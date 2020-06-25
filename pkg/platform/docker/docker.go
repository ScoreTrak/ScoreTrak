package docker

import (
	"archive/tar"
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/L1ghtman2k/ScoreTrak/pkg/config"
	"github.com/L1ghtman2k/ScoreTrak/pkg/platform/util"
	"github.com/L1ghtman2k/ScoreTrak/pkg/platform/worker"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/api/types/swarm"
	"github.com/docker/docker/client"
	"io"
	"os"
	"path/filepath"
)

type Docker struct {
	NetworkName string
	Name        string
	IsSwarm     bool
	Client      *client.Client
	Context     context.Context
}

func NewDocker(config config.StaticConfig) (d *Docker, err error) {
	d = &Docker{NetworkName: config.Platform.Docker.NetworkName, Name: config.Platform.Docker.Name, Context: context.Background()}
	if config.Platform.Use == "swarm" { //https://github.com/openbaton/go-docker-vnfm/blob/8d0a99b48e57d4b94fa14cdb377abe07eaa6c0aa/handler/docker_utils.go#L113
		d.IsSwarm = true
	}
	defaultHeaders := map[string]string{"User-Agent": "engine-api-cli-1.0"}
	d.Client, err = client.NewClient(config.Platform.Docker.Host, "v1.40", nil, defaultHeaders)
	if err != nil {
		return nil, err
	}
	return d, nil
}

func (d *Docker) DeployWorkers(info worker.Info) error {
	err := d.BuildWorkerImage()
	if err != nil {
		return err
	}
	networkName := d.Name + "_" + d.NetworkName
	resp, err := d.CreateWorkerContainer(networkName)
	if err != nil {
		return err
	}
	err = d.UploadConfigToContainer(resp, info)
	if err != nil {
	}
	if d.IsSwarm {
		id, err := d.CommitWorkerContainerToImage(resp)
		if err != nil {
			return err
		}
		_, err = d.CreateService(id, info, networkName)
		if err != nil {
			return nil
		}
	} else {
		err := d.StartContainer(resp)
		if err != nil {
			return err
		}
	}
	return nil
}

func (d *Docker) CommitWorkerContainerToImage(resp container.ContainerCreateCreatedBody) (string, error) {
	id, err := d.Client.ContainerCommit(d.Context, resp.ID, types.ContainerCommitOptions{Reference: "workerSwarm"})
	if err != nil {
		return "", err
	}
	return id.ID, nil
}

func (d *Docker) BuildWorkerImage() error {
	buf := new(bytes.Buffer)
	tw := tar.NewWriter(buf)
	defer tw.Close()
	dockerFile := "deployments/worker/Dockerfile"
	for _, f := range []string{dockerFile, "cmd/worker", "pkg/", "go.mod", "go.sum"} {
		err := util.TarRecurse(f, tw, "test")
		if err != nil {
			return err
		}
	}
	dockerFileTarReader := bytes.NewReader(buf.Bytes())
	imageBuildResponse, err := d.Client.ImageBuild(
		d.Context,
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

func (d *Docker) RemoveWorkers(info worker.Info) error {
	if d.IsSwarm {
		services, err := d.Client.ServiceList(context.Background(), types.ServiceListOptions{})
		if err != nil {
			return err
		}
		for _, service := range services {
			if service.Spec.Name == "worker_"+info.Topic {
				err := d.Client.ServiceRemove(context.Background(), service.ID)
				if err != nil {
					return err
				}
			}
		}
	} else {

	}
	return nil

}

func (d *Docker) CreateService(id string, worker worker.Info, networkName string) (types.ServiceCreateResponse, error) {
	maxAttempts := uint64(1)
	spec := swarm.ServiceSpec{
		Annotations: swarm.Annotations{Name: "worker_" + worker.Topic},
		TaskTemplate: swarm.TaskSpec{
			RestartPolicy: &swarm.RestartPolicy{
				MaxAttempts: &maxAttempts,
				Condition:   swarm.RestartPolicyConditionNone,
			},
			ContainerSpec: swarm.ContainerSpec{
				Image: id,
			},
			Networks: []swarm.NetworkAttachmentConfig{{Target: networkName}},
			Placement: &swarm.Placement{
				Constraints: []string{fmt.Sprintf("node.labels.%s == true", worker.Topic), "node.role != manager"},
			},
		},
	}
	spec.Mode.Global = &swarm.GlobalService{}
	createOptions := types.ServiceCreateOptions{}
	createResponse, err := d.Client.ServiceCreate(context.Background(), spec, createOptions)
	if err != nil {
		return types.ServiceCreateResponse{}, err
	}
	return createResponse, nil
}

func (d *Docker) CreateWorkerContainer(networkName string) (container.ContainerCreateCreatedBody, error) {
	resp, err := d.Client.ContainerCreate(d.Context, &container.Config{
		Image: "worker",
		Tty:   true,
	}, nil, &network.NetworkingConfig{EndpointsConfig: map[string]*network.EndpointSettings{networkName: {}}}, "")
	if err != nil {
		return container.ContainerCreateCreatedBody{}, err
	}
	return resp, nil
}

func (d *Docker) UploadConfigToContainer(resp container.ContainerCreateCreatedBody, info worker.Info) error {
	tmp := filepath.Join(".", "tmp")
	err := os.MkdirAll(tmp, os.ModePerm)
	if err != nil {
		return err
	}

	cnf, err := config.GetConfigCopy()
	if err != nil {
		return err
	}
	if cnf.Queue.Use == "nsq" {
		cnf.Queue.NSQ.Topic = info.Topic
	} else {
		return errors.New("selected queue is not yet supported with platform Docker")
	}
	path := fmt.Sprintf("tmp/config_worker_%s", info.Topic)
	err = config.SaveConfigToYamlFile(path, cnf)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	tw := tar.NewWriter(buf)
	defer tw.Close()
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	defer os.Remove(path)
	s, err := file.Stat()
	if err != nil {
		return err
	}
	tarHeader := &tar.Header{
		Name: "configs/config.yml",
		Size: s.Size(),
		Mode: int64(s.Mode()),
	}
	err = tw.WriteHeader(tarHeader)
	if err != nil {
		return err
	}
	_, err = io.Copy(tw, file)
	if err != nil {
		return err
	}
	dockerFileTarReader := bytes.NewReader(buf.Bytes())
	err = d.Client.CopyToContainer(d.Context, resp.ID, "/go/src/github.com/L1ghtman2k/ScoreTrak", dockerFileTarReader, types.CopyToContainerOptions{AllowOverwriteDirWithFile: true})
	if err != nil {
		return err
	}
	return nil
}

func (d *Docker) StartContainer(resp container.ContainerCreateCreatedBody) error {
	if err := d.Client.ContainerStart(d.Context, resp.ID, types.ContainerStartOptions{}); err != nil {
		return err
	}
	//_, err := d.Client.ContainerWait(d.Context, resp.ID)
	//fmt.Println("But not Here!")
	//if err != nil{
	//	return err
	//}
	//_, err = d.Client.ContainerLogs(d.Context, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
	//if err != nil {
	//	return err
	//}
	return nil
}

func (d *Docker) PullImage() (io.ReadCloser, error) {
	reader, err := d.Client.ImagePull(d.Context, "docker.io/library/alpine", types.ImagePullOptions{})
	if err != nil {
		return nil, err
	}
	return reader, err
}
