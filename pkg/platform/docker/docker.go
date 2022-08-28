package docker

import (
	"archive/tar"
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/ScoreTrak/ScoreTrak/pkg/config"
	"github.com/ScoreTrak/ScoreTrak/pkg/platform/util"
	"github.com/ScoreTrak/ScoreTrak/pkg/platform/worker"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/api/types/swarm"
	"github.com/docker/docker/client"
)

type Docker struct {
	NetworkName string
	Name        string
	IsSwarm     bool
	Client      *client.Client
	Config      config.StaticConfig
}

func NewDocker(cfg config.StaticConfig) (d *Docker, err error) {
	d = &Docker{NetworkName: cfg.Platform.Docker.Network, Name: cfg.Platform.Docker.Name, Config: cfg}
	if cfg.Platform.Use == "swarm" { //https://github.com/openbaton/go-docker-vnfm/blob/8d0a99b48e57d4b94fa14cdb377abe07eaa6c0aa/handler/docker_utils.go#L113
		d.IsSwarm = true
	}
	defaultHeaders := map[string]string{"User-Agent": "engine-api-cli-1.0"}
	d.Client, err = client.NewClient(cfg.Platform.Docker.Host, "v1.40", nil, defaultHeaders)
	if err != nil {
		return nil, err
	}
	return d, nil
}

// func (d *Docker) GetWorkerServiceStatus(info worker.Info)(status string, err error){
// 	s, err := d.GetServiceByName("worker_"+info.Topic)
// 	if err != nil{
// 		return "", err
// 	}
// }

func (d *Docker) GetWorkerContainerStatus(ctx context.Context, info worker.Info) (status string, err error) {
	ctr, err := d.GetContainerByName(ctx, "worker_"+info.Topic)
	if err != nil {
		return "", err
	}
	return ctr.Status, nil
}

var ErrLabelIsEmpty = errors.New("label should not be empty when creating a check_service on swarm platform")

func (d *Docker) DeployWorkers(ctx context.Context, info worker.Info) (err error) {
	networkName := d.Name + "_" + d.NetworkName
	if !d.IsSwarm {
		resp, err := d.CreateWorkerContainer(ctx, networkName, info)
		if err != nil {
			return err
		}
		err = d.StartContainer(ctx, resp)
		if err != nil {
			return err
		}
	} else {
		if info.Label == "" {
			return ErrLabelIsEmpty
		}
		_, err := d.CreateService(ctx, info, networkName)
		if err != nil {
			return err
		}
	}

	return nil
}

func (d *Docker) RemoveWorkers(ctx context.Context, info worker.Info) error {
	if d.IsSwarm {
		s, err := d.GetServiceByName(ctx, "worker_"+info.Topic)
		if err != nil {
			return err
		}
		return d.Client.ServiceRemove(ctx, s.ID)
	}
	ctr, err := d.GetContainerByName(ctx, "worker_"+info.Topic)
	if err != nil {
		return err
	}
	return d.Client.ContainerRemove(ctx, ctr.ID, types.ContainerRemoveOptions{Force: true})
}

var ErrCheckServiceMissing = errors.New("unable to find check_service. The workers might have already been removed")

func (d *Docker) GetServiceByName(ctx context.Context, serviceName string) (swarm.Service, error) {
	services, err := d.Client.ServiceList(ctx, types.ServiceListOptions{})
	if err != nil {
		return swarm.Service{}, err
	}
	for _, service := range services {
		if strings.Contains(service.Spec.Name, serviceName) {
			return service, nil
		}
	}

	return swarm.Service{}, ErrCheckServiceMissing
}

var ErrContainerNotFound = errors.New("container not found. The worker might have already been removed")

func (d *Docker) GetContainerByName(ctx context.Context, containerName string) (types.Container, error) {
	containers, err := d.Client.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		return types.Container{}, err
	}
	for _, ctr := range containers {
		for _, name := range ctr.Names {
			if strings.Contains(name, containerName) {
				return ctr, nil
			}
		}
	}
	return types.Container{}, ErrContainerNotFound
}

func (d *Docker) CommitWorkerContainerToImage(ctx context.Context, resp container.ContainerCreateCreatedBody, info worker.Info) (string, error) {
	id, err := d.Client.ContainerCommit(ctx, resp.ID, types.ContainerCommitOptions{Reference: "worker_" + info.Topic})
	if err != nil {
		return "", err
	}
	return id.ID, nil
}

func (d *Docker) BuildWorkerImage(ctx context.Context) error {
	buf := new(bytes.Buffer)
	tw := tar.NewWriter(buf)
	defer tw.Close()
	dockerFile := "deployments/Dockerfile"
	for _, f := range []string{dockerFile, "cmd/worker", "pkg/", "go.mod", "go.sum"} {
		err := util.TarRecurse(f, tw, "test")
		if err != nil {
			return err
		}
	}
	dockerFileTarReader := bytes.NewReader(buf.Bytes())
	imageBuildResponse, err := d.Client.ImageBuild(
		ctx,
		dockerFileTarReader,
		types.ImageBuildOptions{
			Context:    dockerFileTarReader,
			Dockerfile: dockerFile,
			Remove:     true,
			Tags:       []string{"worker"},
		})
	if err != nil {
		return err
	}
	defer imageBuildResponse.Body.Close()
	_, err = io.Copy(os.Stdout, imageBuildResponse.Body)
	if err != nil {
		return err
	}
	return nil
}

func (d *Docker) CreateService(ctx context.Context, info worker.Info, networkName string) (types.ServiceCreateResponse, error) {
	encWorkerCfg, err := util.GenerateEncodedWorkerCfg(d.Config, info)
	if err != nil {
		return types.ServiceCreateResponse{}, err
	}
	maxAttempts := uint64(1)
	spec := swarm.ServiceSpec{
		Annotations: swarm.Annotations{Name: "worker_" + info.Topic},
		Mode:        swarm.ServiceMode{Global: &swarm.GlobalService{}},
		TaskTemplate: swarm.TaskSpec{
			RestartPolicy: &swarm.RestartPolicy{
				MaxAttempts: &maxAttempts,
				Condition:   swarm.RestartPolicyConditionOnFailure,
			},
			ContainerSpec: swarm.ContainerSpec{
				Image: util.Image,
				Args:  []string{"worker", "--encoded-config", encWorkerCfg},
			},
			Networks: []swarm.NetworkAttachmentConfig{{Target: networkName}},
			Placement: &swarm.Placement{
				Constraints: []string{fmt.Sprintf("node.labels.%s == true", info.Label)},
			},
		},
	}
	spec.Mode.Global = &swarm.GlobalService{}
	createOptions := types.ServiceCreateOptions{}
	createResponse, err := d.Client.ServiceCreate(ctx, spec, createOptions)
	if err != nil {
		return types.ServiceCreateResponse{}, err
	}
	return createResponse, nil
}

func (d *Docker) CreateWorkerContainer(ctx context.Context, networkName string, info worker.Info) (container.ContainerCreateCreatedBody, error) {
	encWorkerCfg, err := util.GenerateEncodedWorkerCfg(d.Config, info)
	if err != nil {
		return container.ContainerCreateCreatedBody{}, err
	}
	resp, err := d.Client.ContainerCreate(ctx, &container.Config{
		Image: util.Image,
		Tty:   true,
		Cmd:   []string{"worker", "--encoded-config", encWorkerCfg},
	}, nil, &network.NetworkingConfig{EndpointsConfig: map[string]*network.EndpointSettings{networkName: {}}}, "worker_"+info.Topic)
	if err != nil {
		return container.ContainerCreateCreatedBody{}, err
	}
	return resp, nil
}

func (d *Docker) UploadConfigToContainer(ctx context.Context, resp container.ContainerCreateCreatedBody, path string) error {
	buf := new(bytes.Buffer)
	writer := tar.NewWriter(buf)
	defer writer.Close()
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	defer os.Remove(path)
	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}
	tarHeader := &tar.Header{
		Name: "configs/config.yml",
		Size: fileInfo.Size(),
		Mode: int64(fileInfo.Mode()),
	}
	err = writer.WriteHeader(tarHeader)
	if err != nil {
		return err
	}
	_, err = io.Copy(writer, file)
	if err != nil {
		return err
	}
	dockerFileTarReader := bytes.NewReader(buf.Bytes())
	err = d.Client.CopyToContainer(ctx, resp.ID, "/go/src/github.com/ScoreTrak/ScoreTrak", dockerFileTarReader, types.CopyToContainerOptions{AllowOverwriteDirWithFile: true})
	if err != nil {
		return err
	}
	return nil
}

func (d *Docker) StartContainer(ctx context.Context, resp container.ContainerCreateCreatedBody) error {
	if err := d.Client.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		return err
	}
	// _, err := d.Client.ContainerWait(ctx, resp.ID)
	// fmt.Println("But not Here!")
	// if err != nil{
	// 	return err
	// }
	// _, err = d.Client.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
	// if err != nil {
	// 	return err
	// }
	return nil
}

func (d *Docker) PullImage(ctx context.Context) error {
	reader, err := d.Client.ImagePull(ctx, util.Image, types.ImagePullOptions{})
	if err != nil {
		return err
	}
	_, err = io.Copy(os.Stdout, reader)
	if err != nil {
		return err
	}
	err = reader.Close()
	if err != nil {
		return err
	}
	return nil
}
