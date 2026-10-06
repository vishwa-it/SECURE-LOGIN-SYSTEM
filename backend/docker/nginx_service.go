package docker

import (
	"context"
	"io"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

const (
	nginxContainerName = "secure-login-nginx"
	nginxImage         = "nginx:latest"
	nginxHostPort      = "8081"
)

func StartNginx() error {
	ctx := context.Background()

	cli, err := client.NewClientWithOpts(
		client.FromEnv,
		client.WithAPIVersionNegotiation(),
	)
	if err != nil {
		return err
	}
	defer cli.Close()

	containers, err := cli.ContainerList(ctx, container.ListOptions{All: true})
	if err != nil {
		return err
	}

	for _, c := range containers {
		for _, name := range c.Names {
			if name != "/"+nginxContainerName {
				continue
			}

			if c.State == "running" {
				return nil
			}

			return cli.ContainerStart(ctx, c.ID, container.StartOptions{})
		}
	}

	if err := ensureImage(ctx, cli); err != nil {
		return err
	}

	containerConfig := &container.Config{
		Image: nginxImage,
	}

	hostConfig := &container.HostConfig{
		PortBindings: nat.PortMap{
			"80/tcp": []nat.PortBinding{
				{
					HostIP:   "0.0.0.0",
					HostPort: nginxHostPort,
				},
			},
		},
	}

	resp, err := cli.ContainerCreate(
		ctx,
		containerConfig,
		hostConfig,
		nil,
		nil,
		nginxContainerName,
	)
	if err != nil {
		return err
	}

	return cli.ContainerStart(ctx, resp.ID, container.StartOptions{})
}

func ensureImage(ctx context.Context, cli *client.Client) error {
	if _, _, err := cli.ImageInspectWithRaw(ctx, nginxImage); err == nil {
		return nil
	}

	reader, err := cli.ImagePull(ctx, nginxImage, image.PullOptions{})
	if err != nil {
		return err
	}
	defer reader.Close()

	_, err = io.Copy(io.Discard, reader)
	return err
}
