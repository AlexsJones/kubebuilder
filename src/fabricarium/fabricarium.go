package fabricarium

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path"

	"github.com/AlexsJones/kubebuilder/src/data"
	"github.com/AlexsJones/kubebuilder/src/fabricarium/vcs"
	"github.com/AlexsJones/kubebuilder/src/log"
	shortid "github.com/ventu-io/go-shortid"
	validator "gopkg.in/go-playground/validator.v9"
)

//Fabricarium object that maintains a small amount of state
type Fabricarium struct {
	Configuration *Configuration
}

//Mount ...
type Mount struct {
	Path string
}

//Configuration holds configuration information
type Configuration struct {
	MountInformation *Mount
}

//NewFabricarium creates the builder that receivees YAML build scripts
func NewFabricarium(conf *Configuration) *Fabricarium {

	return &Fabricarium{Configuration: conf}

}

/*
vcs:
  type: git
  sshPath: git@github.com/AlexsJones/jnxlibc.git
  checkoutArgs: ""
build:
  commands:
    - cmake .
    - make
  docker:
    imageNameSuffix: vTEST
    buildArgs:
k8s:
  deployment: ./k8s/deployment.yaml
  imagePlaceholderReplacement: latest
*/
//Process does the checkout and construction
func (f *Fabricarium) Process(build *data.BuildDefinition) {

	logger.GetInstance().Log(fmt.Sprintf("Processing build: %v", build))

	validate := validator.New()
	err := validate.Struct(build)
	if err != nil {

		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like this.
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return
		}
	}
	sid, err := shortid.New(1, shortid.DefaultABC, 2342)
	if err != nil {
		panic(err)
	}
	//Create the mount directory if it doesn't exist
	if _, err := os.Stat(f.Configuration.MountInformation.Path); os.IsNotExist(err) {
		os.Mkdir(f.Configuration.MountInformation.Path, 0700)
	}
	//Create the sub directory for the repo if that doesnt exist
	if _, err := os.Stat(path.Join(f.Configuration.MountInformation.Path, build.VCS.Name)); os.IsNotExist(err) {
		os.Mkdir(path.Join(f.Configuration.MountInformation.Path, build.VCS.Name), 0700)
	}

	dynamicBuildPath := path.Join(f.Configuration.MountInformation.Path, build.VCS.Name, sid.MustGenerate())

	logger.GetInstance().Log(fmt.Sprintf("Created new dynamic build path %s", dynamicBuildPath))
	//VCS
	if err := f.processVCS(dynamicBuildPath, build); err != nil {
		logger.GetInstance().Log(err.Error())
	}
	//Build
	if err := f.processBuild(dynamicBuildPath, build); err != nil {
		logger.GetInstance().Log(err.Error())
	}
	//K8s
	if err := f.processK8s(dynamicBuildPath, build); err != nil {
		logger.GetInstance().Fatal(err.Error())
		return
	}
}

func (f *Fabricarium) processVCS(dynamicBuildPath string, build *data.BuildDefinition) error {
	logger.GetInstance().Log("-------------------------Processing VCS-------------------------")

	switch build.VCS.Type {
	case "git":
		g := vcs.GitVCS{}
		if err := vcs.Fetch(g, dynamicBuildPath, build); err != nil {
			return err
		}
	default:
		return errors.New("Unknown build type cannot continue")
	}

	return nil
}

func (f *Fabricarium) processBuild(dynamicBuildPath string, build *data.BuildDefinition) error {
	logger.GetInstance().Log("-------------------------Processing Build-------------------------")

	//Run build commands

	logger.GetInstance().Log(fmt.Sprintf("/bin/sh %s", build.Build.Commands))
	cmd := exec.Command("/bin/bash", "-c", build.Build.Commands)
	cmd.Dir = dynamicBuildPath
	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating StdoutPipe for Cmd", err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(cmdReader)
	go func() {
		for scanner.Scan() {
			logger.GetInstance().Log(fmt.Sprintf(scanner.Text()))
		}
	}()

	err = cmd.Start()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error starting Cmd", err)
		return err
	}

	err = cmd.Wait()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error waiting for Cmd", err)
		return err
	}
	return nil
}

func (f *Fabricarium) processK8s(dynamicBuildPath string, build *data.BuildDefinition) error {
	logger.GetInstance().Log("-------------------------Processing K8s-------------------------")

	//TODO

	return nil
}
