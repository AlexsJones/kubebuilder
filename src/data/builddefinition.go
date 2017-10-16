package data

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
	Tag       string `yaml:"tag"`
	Buildargs string `yaml:"buildArgs"`
}

//Build ...
type Build struct {
	Commands string `yaml:"commands" validate:"required"`
	Docker   Docker `yaml:"docker"`
}

//K8s ...
type Kubernetes struct {
	Namespace                   string `yaml:"namespace" validate:"required"`
	Deployment                  string `yaml:"deployment" validate:"required"`
	ImagePlaceholderReplacement string `yaml:"imagePlaceholderReplacement" validate:"required"`
}

//BuildDefinition ...
type BuildDefinition struct {
	VCS        VCS        `yaml:"vcs" validate:"required"`
	Build      Build      `yaml:"build" validate:"required"`
	Kubernetes Kubernetes `yaml:"kubernetes" validate:"required"`
}
