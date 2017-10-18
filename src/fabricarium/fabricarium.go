package fabricarium

import (
	"errors"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/AlexsJones/kubebuilder/src/config"
	"github.com/AlexsJones/kubebuilder/src/data"
	"github.com/AlexsJones/kubebuilder/src/fabricarium/container"
	"github.com/AlexsJones/kubebuilder/src/fabricarium/platform"
	"github.com/AlexsJones/kubebuilder/src/fabricarium/vcs"
	"github.com/AlexsJones/kubebuilder/src/log"
	sh "github.com/AlexsJones/kubebuilder/src/shell"
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
	MountInformation         *Mount
	ApplicationConfiguration *config.Configuration
}

//NewFabricarium creates the builder that receivees YAML build scripts
func NewFabricarium(conf *Configuration) *Fabricarium {

	return &Fabricarium{Configuration: conf}

}

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

	if f.Configuration.ApplicationConfiguration.KubeBuilderConfiguration.BypassVCS {
		logger.GetInstance().Info("Bypassing VCS")
	} else {
		if err := f.processVCS(dynamicBuildPath, build); err != nil {
			logger.GetInstance().Log(err.Error())
		}
	}
	//Build
	if f.Configuration.ApplicationConfiguration.KubeBuilderConfiguration.BypassBuild {
		logger.GetInstance().Info("Bypassing Build")
	} else {
		if err := f.processBuild(dynamicBuildPath, build); err != nil {
			logger.GetInstance().Log(err.Error())
		}
	}
	//K8s
	if f.Configuration.ApplicationConfiguration.KubeBuilderConfiguration.BypassKubernetes {
		logger.GetInstance().Info("Bypassing Kubernetes")
	} else {
		if err := f.processK8s(dynamicBuildPath, build); err != nil {
			logger.GetInstance().Fatal(err.Error())
			return
		}
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

	dockerClient := container.NewDocker()

	containerID, err := dockerClient.CreateContainerName(build.Build.Docker.ContainerID, build.Build.Docker.Tag, build.Build.Docker.TagReplacementValue)
	if err != nil {
		return err
	}
	//Lets modify the build commands if it references our replacement value
	buildCommands := strings.Replace(build.Build.Commands, build.Build.Docker.TagReplacementValue, build.Build.Docker.Tag, -1)

	if err = sh.RunCommand(dynamicBuildPath, buildCommands); err != nil {
		return err
	}

	logger.GetInstance().Log(fmt.Sprintf("Using docker container %s\n", containerID))
	if err != nil {
		return err
	}
	//Verify build
	if ok, err := container.Exists(dockerClient, containerID); !ok {
		return err
	}

	//VerifyAge
	if ok, err := container.YoungerThan(dockerClient, containerID, 5); !ok {
		return err
	}

	//tag command
	remoteURL := strings.Replace(build.Build.Docker.Buildargs.URL, build.Build.Docker.TagReplacementValue, build.Build.Docker.Tag, -1)
	logger.GetInstance().Log(fmt.Sprintf("Using remote URL %s", remoteURL))

	tagCommand := fmt.Sprintf("gcloud docker -- tag %s %s", containerID, remoteURL)
	//Deploy build
	if err = sh.RunCommand(dynamicBuildPath, tagCommand); err != nil {
		return err
	}
	logger.GetInstance().Log(tagCommand)

	pushCommand := fmt.Sprintf("gcloud docker -- push %s", remoteURL)
	if err = sh.RunCommand(dynamicBuildPath, pushCommand); err != nil {
		return err
	}
	logger.GetInstance().Log(pushCommand)

	return nil
}

func (f *Fabricarium) processK8s(dynamicBuildPath string, build *data.BuildDefinition) error {
	logger.GetInstance().Log("-------------------------Processing K8s-------------------------")

	//TODO

	//namespace
	_, _ = platform.NewKubernetes(f.Configuration.ApplicationConfiguration.KubernetesConfiguration.InCluster)

	//Deployment

	return nil
}
