package plugins

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/registry"
	"testing"
)

func TestDocker(t *testing.T) {
	dm, err := NewDockerManager()
	if err != nil {
		panic(err)
	}
	okBody, err := dm.cli.RegistryLogin(context.Background(), registry.AuthConfig{
		Username:      "go-developer",
		Password:      "uq123456",
		ServerAddress: "nexus.uqvalley.top:8082",
	})
	if err != nil {
		panic(err)
	}
	t.Logf("%+v\n", okBody)

	//dm.cli.ImagePull()

	list, err := dm.cli.ImageList(context.Background(), types.ImageListOptions{All: true})
	if err != nil {
		panic(err)
	}
	for _, item := range list {
		t.Logf("%+v\n", item)
	}

	info, err := dm.cli.Info(context.Background())
	if err != nil {
		panic(err)
	}
	for _, cfg := range info.RegistryConfig.IndexConfigs {
		t.Logf("%+v\n", cfg)
	}

	//containerConfig := container.Config{
	//	Image: "nginx",
	//}
	//resp, err := dm.cli.ContainerCreate(context.Background(), &containerConfig, &container.HostConfig{}, &network.NetworkingConfig{}, nil, "nginx_test")
	//if err != nil {
	//	panic(err)
	//}
	//t.Log(resp)
}
