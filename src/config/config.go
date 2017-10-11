package config

import (
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

//Configuration representation of yaml
type Configuration struct {
	GCPConfiguration GCPConfiguration `yaml:"gcp"`
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
