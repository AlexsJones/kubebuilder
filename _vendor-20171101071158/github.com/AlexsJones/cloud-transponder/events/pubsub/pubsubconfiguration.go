package gcloud

import (
	"errors"
)

//PubSubConfiguration ...
type PubSubConfiguration struct {
	Topic              string
	ConnectionString   string
	SubscriptionString string
}

//NewPubSubConfiguration ...
func NewPubSubConfiguration() *PubSubConfiguration {

	return &PubSubConfiguration{}
}

//GetTopic ...
func (p *PubSubConfiguration) GetTopic() (string, error) {
	if p.Topic == "" {
		return "", errors.New("No topic set")
	}
	return p.Topic, nil
}

//GetConnectionString ...
func (p *PubSubConfiguration) GetConnectionString() (string, error) {
	if p.ConnectionString == "" {
		return "", errors.New("No connection string set")
	}
	return p.ConnectionString, nil
}

//GetSubscriptionString ...
func (p *PubSubConfiguration) GetSubscriptionString() (string, error) {
	if p.SubscriptionString == "" {
		return "", errors.New("No subscription string set")
	}
	return p.SubscriptionString, nil
}
