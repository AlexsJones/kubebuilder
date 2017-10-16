package container

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/AlexsJones/kubebuilder/src/log"
	"github.com/docker/docker/client"
)

type Docker struct {
	cli *client.Client
	ctx *context.Context
}

func NewDocker() *Docker {
	ctx := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		logger.GetInstance().Log(err.Error())
		return nil
	}
	return &Docker{cli: cli, ctx: &ctx}
}
func (d *Docker) Exists(containerID string) (bool, error) {

	cont, err := d.cli.ImageHistory(*d.ctx, containerID)
	if err != nil {
		logger.GetInstance().Log(err.Error())
		return false, err
	}
	if len(cont) < 1 {
		return false, errors.New("Nothing returned in image list")
	}
	return true, nil
}

func (d *Docker) YoungerThan(containerID string, age float64) (bool, error) {
	cont, err := d.cli.ImageHistory(*d.ctx, containerID)
	if err != nil {
		logger.GetInstance().Log(err.Error())
		return false, err
	}
	if len(cont) < 1 {
		return false, errors.New("Nothing returned in image list")
	}
	t := time.Unix(0, cont[0].Created)
	elapsed := time.Since(t)

	if elapsed.Minutes() > age {
		logger.GetInstance().Log(fmt.Sprintf("Image is older than %d minutes", age))
		return false, errors.New(fmt.Sprintf("Image is older than %d minutes", age))
	}
	logger.GetInstance().Log("Image looks good")
	return true, nil
}
