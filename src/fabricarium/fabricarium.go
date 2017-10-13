package fabricarium

//Fabricarium object that maintains a small amount of state
type Fabricarium struct {
	Configuration *Configuration
}

//Configuration holds configuration information
type Configuration struct {
	Mount struct {
		Path string
	}
}

//NewFabricarium creates the builder that receivees YAML build scripts
func NewFabricarium(conf *Configuration) *Fabricarium {

	return &Fabricarium{Configuration: conf}

}
