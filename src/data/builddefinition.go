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
	Type         string `yaml:"type"`
	SSHPath      string `yaml:"sshPath"`
	CheckoutArgs string `yaml:"checkoutArgs"`
}

//Docker ...
type Docker struct {
	ImageNameSuffix string `yaml:"imageNameSuffix"`
	Buildargs       string `yaml:"buildArgs"`
}

//Build ...
type Build struct {
	Commands []string `yaml:"commands"`
	Docker   Docker   `yaml:"docker"`
}

//K8s ...
type K8s struct {
	Deployment                  string `yaml:"deployment"`
	ImagePlaceholderReplacement string `yaml:"imagePlaceholderReplacement"`
}

//BuildDefinition ...
type BuildDefinition struct {
	VCS   VCS   `yaml:"vcs"`
	Build Build `yaml:"build"`
	K8s   K8s   `yaml:"k8s"`
}
