package plugins

import "github.com/docker/docker/client"

type DockerManager struct {
	cli *client.Client
}

func NewDockerManager() (dm *DockerManager, err error) {
	dm = &DockerManager{}
	dm.cli, err = client.NewClientWithOpts()
	if err != nil {
		return nil, err
	}
	return dm, err
}
