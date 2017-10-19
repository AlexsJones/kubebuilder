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
	bypassVCS: true
	bypassBuild: true
	bypassKubernetes: false
gcp:
  connectionString: beamery-trials
  topic: cadium
  subscriptionString: cadium-sub
*/
//Configuration representation of yaml
type Configuration struct {
	PubSubConfiguration      PubSubConfiguration      `yaml:"pubsub" validate:"required"`
	KubeBuilderConfiguration KubeBuilderConfiguration `yaml:"kubebuilder" validate:"required"`
	KubernetesConfiguration  KubernetesConfiguration  `yaml:"kubernetes" validate:"required"`
}

//KubeBuilderConfiguration applicatoin configuration
type KubeBuilderConfiguration struct {
	Mount            string `yaml:"mount" validate:"required"`
	Drainqueue       bool   `yaml:"drainqueue" validate:"required"`
	BypassVCS        bool   `yaml:"bypassVCS" validate:"required"`
	BypassBuild      bool   `yaml:"bypassBuild" validate:"required"`
	BypassKubernetes bool   `yaml:"bypassKubernetes" validate:"required"`
}

//PubSubConfiguration object
type PubSubConfiguration struct {
	ConnectionString   string `yaml:"connectionString"`
	Topic              string `yaml:"topic"`
	SubscriptionString string `yaml:"subscriptionString"`
}

//KubernetesConfiguration ...
type KubernetesConfiguration struct {
	InCluster bool   `yaml:"incluster" validate:"required"`
	MasterURL string `yaml:"masterurl" validate:"required"`
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
