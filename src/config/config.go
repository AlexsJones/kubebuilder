package config

import (
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

/*
kubebuilder:
  mount: ./mount
  purgepubsub: true
gcp:
  connectionString: beamery-trials
  topic: cadium
  subscriptionString: cadium-sub
*/
//Configuration representation of yaml
type Configuration struct {
	GCPConfiguration         GCPConfiguration         `yaml:"gcp"`
	KubeBuilderConfiguration KubeBuilderConfiguration `yaml:"kubebuilder" validate:"required"`
}

//KubeBuilderConfiguration applicatoin configuration
type KubeBuilderConfiguration struct {
	Mount      string `yaml:"mount" validate:"required"`
	Drainqueue bool   `yaml:drainqueue validate:"required"`
}

//GCPConfiguration object
type GCPConfiguration struct {
	ConnectionString   string `yaml:"connectionString"`
	Topic              string `yaml:"topic"`
	SubscriptionString string `yaml:"subscriptionString"`
}

//LoadConfiguration ...
func LoadConfiguration(filepath string) (*Configuration, error) {
	bytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	m := &Configuration{}
	err = yaml.Unmarshal(bytes, m)
	if err != nil {
		log.Print("yaml: ", err)
		return nil, err
	}
	return m, err
}
