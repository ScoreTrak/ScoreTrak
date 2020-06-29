package docker

import (
	"archive/tar"
	"bytes"
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/L1ghtman2k/ScoreTrak/pkg/config"
	"github.com/L1ghtman2k/ScoreTrak/pkg/logger"
	"github.com/L1ghtman2k/ScoreTrak/pkg/platform/util"
	"github.com/L1ghtman2k/ScoreTrak/pkg/platform/worker"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/api/types/swarm"
	"github.com/docker/docker/client"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Docker struct {
	l           logger.LogInfoFormat
	NetworkName string
	Name        string
	IsSwarm     bool
	Client      *client.Client
	Context     context.Context
}

func NewDocker(cnf config.StaticConfig, l logger.LogInfoFormat) (d *Docker, err error) {
	d = &Docker{NetworkName: cnf.Platform.Docker.Network, Name: cnf.Platform.Docker.Name, Context: context.Background(), l: l}
	if cnf.Platform.Use == "swarm" { //https://github.com/openbaton/go-docker-vnfm/blob/8d0a99b48e57d4b94fa14cdb377abe07eaa6c0aa/handler/docker_utils.go#L113
		d.IsSwarm = true
	}
	defaultHeaders := map[string]string{"User-Agent": "engine-api-cli-1.0"}
	d.Client, err = client.NewClient(cnf.Platform.Docker.Host, "v1.40", nil, defaultHeaders)
	if err != nil {
		return nil, err
	}
	return d, nil
}

func (d *Docker) DeployWorkers(info worker.Info) (err error) {
	networkName := d.Name + "_" + d.NetworkName
	tmp := filepath.Join(".", "tmp")
	err = os.MkdirAll(tmp, os.ModePerm)
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
	if !d.IsSwarm {
		err = d.PullImage()
		if err != nil {
			return err
		}
		resp, err := d.CreateWorkerContainer(networkName)
		if err != nil {
			d.l.Error(err)
			return err
		}
		err = d.UploadConfigToContainer(resp, path)
		if err != nil {
			d.l.Error(err)
			return err
		}
		err = d.StartContainer(resp)
		if err != nil {
			d.l.Error(err)
			return err
		}
	} else {
		_, err := d.CreateService(info, networkName, path)
		if err != nil {
			d.l.Error(err)
			return nil
		}
	}

	return nil
}

func (d *Docker) CommitWorkerContainerToImage(resp container.ContainerCreateCreatedBody, info worker.Info) (string, error) {
	id, err := d.Client.ContainerCommit(d.Context, resp.ID, types.ContainerCommitOptions{Reference: "worker_" + info.Topic})
	if err != nil {
		return "", err
	}
	return id.ID, nil
}

func (d *Docker) BuildWorkerImage() error {
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

func (d *Docker) CreateService(info worker.Info, networkName string, configPath string) (types.ServiceCreateResponse, error) {
	content, err := ioutil.ReadFile(configPath)
	if err != nil {
		return types.ServiceCreateResponse{}, err
	}
	cEnc := base64.StdEncoding.EncodeToString(content)
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
				Image:   "l1ghtman/scoretrak:latest",
				Command: []string{"./worker", "-encoded-config", cEnc},
			},
			Networks: []swarm.NetworkAttachmentConfig{{Target: networkName}},
			Placement: &swarm.Placement{
				Constraints: []string{fmt.Sprintf("node.labels.%s == true", info.Topic)},
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
		Image: "l1ghtman/scoretrak:latest",
		Tty:   true,
	}, nil, &network.NetworkingConfig{EndpointsConfig: map[string]*network.EndpointSettings{networkName: {}}}, "")
	if err != nil {
		return container.ContainerCreateCreatedBody{}, err
	}
	return resp, nil
}

func (d *Docker) UploadConfigToContainer(resp container.ContainerCreateCreatedBody, path string) error {

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

func (d *Docker) PullImage() error {
	reader, err := d.Client.ImagePull(d.Context, "docker.io/l1ghtman/scoretrak:latest", types.ImagePullOptions{})
	if err != nil {
		return err
	}
	io.Copy(os.Stdout, reader)
	reader.Close()
	return nil
}
