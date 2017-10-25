package data

/*
vcs:
  type: git
  name: api-core
  branch: master
  path: git@github.com:SeedJobs/api-core.git
  checkoutArgs: ""
build:
  commands: gcloud docker -- build --no-cache=true -t api-core:TAGNAME .
  docker:
    tag: latest
    tagReplacementValue: TAGNAME
    containerID: api-core:TAGNAME
    buildArgs:
      remote:
      url: us.gcr.io/beamery-trials/api-core:TAGNAME
kubernetes:
  namespace: alex
  deploymentProject: git@github.com:SeedJobs/devops-kubernetes-beamery.git
  commands: ""
  deployment: ./deployment/api-core/deployment.yaml
  imagePlaceholderReplacement: latest
*/

//VCS ...
type VCS struct {
	Type         string `yaml:"type" validate:"required"`
	Path         string `yaml:"path" validate:"required"`
	Name         string `yaml:"name" validate:"required"`
	Branch       string `yaml:"branch" validate:"required"`
	CheckoutArgs string `yaml:"checkoutArgs"`
}

//Docker ...
type Docker struct {
	Tag                 string          `yaml:"tag"`
	TagReplacementValue string          `yaml:"tagReplacementValue"`
	Buildargs           DockerBuildArgs `yaml:"buildArgs"`
	ContainerID         string          `yaml:"containerID"`
}

//DockerBuildArgs ...
type DockerBuildArgs struct {
	Remote string `yaml:"remote"`
	URL    string `yaml:"url"`
}

//Build ...
type Build struct {
	Commands string `yaml:"commands" validate:"required"`
	Docker   Docker `yaml:"docker"`
}

//Kubernetes ...
type Kubernetes struct {
	Namespace                   string `yaml:"namespace" validate:"required"`
	Deployment                  string `yaml:"deployment" validate:"required"`
	PreDeployCommands           string `yaml:"preDeployCommands"`
	PostDeployCommands          string `yaml:"postDeployCommands"`
	DeploymentProjectProtocol   string `yaml:"deploymentProjectProtocol"`
	DeploymentProject           string `yaml:"deploymentProject"`
	ImagePlaceholderReplacement string `yaml:"imagePlaceholderReplacement" validate:"required"`
}

//BuildDefinition ...
type BuildDefinition struct {
	VCS        VCS        `yaml:"vcs" validate:"required"`
	Build      Build      `yaml:"build" validate:"required"`
	Kubernetes Kubernetes `yaml:"kubernetes" validate:"required"`
}
