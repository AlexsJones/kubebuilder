package fabricarium

import "github.com/AlexsJones/kubebuilder/src/data"

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

func (f *Fabricarium) ProcessBuild(build *data.BuildDefinition) {

}
