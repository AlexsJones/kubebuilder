package gcloud

import (
	"cloud.google.com/go/pubsub"
)

//PubSubMessage definition
type PubSubMessage struct {
	Message *pubsub.Message
}

//Ack a reply
func (p *PubSubMessage) Ack() {
	p.Message.Ack()
}

//Nack a reply
func (p *PubSubMessage) Nack() {
	p.Message.Nack()
}

//GetRaw message
func (p *PubSubMessage) GetRaw() []byte {
	return p.Message.Data
}
