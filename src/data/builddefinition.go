package data

/*
vcs:
  type: git
  name: api-core
  branch: master
  path: git@github.com:BLAH/api-example.git
  checkoutArgs: ""
build:
  commands: gcloud docker -- build --no-cache=true -t api-example:TAGNAME .
  docker:
    tag: latest
    tagReplacementValue: TAGNAME
    containerID: api-example:TAGNAME
    buildArgs:
      remote:
      url: us.gcr.io/beamery-trials/api-example:TAGNAME
kubernetes:
  masterurl:
  namespace: alex
  yaml: |
    "apiVersion: extensions/v1beta1
    kind: Deployment
    metadata:
      name: THISISBROKEN
      namespace: THISCANTBETRUE"
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
	Buildargs   DockerBuildArgs `yaml:"buildArgs"`
	ContainerID string          `yaml:"containerID"`
}

//DockerBuildArgs ...
type DockerBuildArgs struct {
	URL string `yaml:"url"`
}

//Build ...
type Build struct {
	Commands string `yaml:"commands" validate:"required"`
	Docker   Docker `yaml:"docker"`
}

//Kubernetes ...
type Kubernetes struct {
	Namespace  string `yaml:"namespace" validate:"required"`
	Deployment string `yaml:"deployment" validate:"required"`
	Service    string `yaml:"service" validate:"required"`
	Ingress    string `yaml:"ingress" validate:"required"`
	Secret     string `yaml:"secret" validate:"required"`
}

//BuildDefinition ...
type BuildDefinition struct {
	VCS        VCS        `yaml:"vcs" validate:"required"`
	Build      Build      `yaml:"build" validate:"required"`
	Kubernetes Kubernetes `yaml:"kubernetes" validate:"required"`
}
