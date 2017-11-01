package container

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/AlexsJones/kubebuilder/src/log"
	"github.com/araddon/dateparse"
	"github.com/docker/docker/client"
)

//Docker ...
type Docker struct {
	cli *client.Client
	ctx *context.Context
}

//NewDocker ...
func NewDocker() *Docker {
	ctx := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		logger.GetInstance().Log(err.Error())
		return nil
	}
	return &Docker{cli: cli, ctx: &ctx}
}

//Exists checks docker images for being real
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

//YoungerThan a time value vs image
func (d *Docker) YoungerThan(containerID string, age float64) (bool, error) {
	cont, _, err := d.cli.ImageInspectWithRaw(*d.ctx, containerID)

	if err != nil {
		logger.GetInstance().Log(err.Error())
		return false, err
	}
	t, err := dateparse.ParseAny(cont.Created)
	elapsed := time.Since(t)
	fmt.Println(t)

	if elapsed.Minutes() > age {
		logger.GetInstance().Log(fmt.Sprintf("Image is older than %v minutes", age))
		return false, fmt.Errorf("Image is older than %v minutes", age)
	}
	logger.GetInstance().Log("Image looks good")
	return true, nil
}
