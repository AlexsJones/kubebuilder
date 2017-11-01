package gcloud

import (
	"errors"
	"log"

	"golang.org/x/net/context"

	"cloud.google.com/go/pubsub"
	event "github.com/AlexsJones/cloud-transponder/events"
)

//PubSub type definitions
type PubSub struct {
	Topic           string
	Project         string
	Subscription    string
	TopicRef        *pubsub.Topic
	ClientRef       *pubsub.Client
	SubscriptionRef *pubsub.Subscription
	Cancel          func()
	Ctx             context.Context
}

//NewPubSub creates a PubSub
func NewPubSub() *PubSub {
	return &PubSub{}
}

//Connect ...
func (p *PubSub) Connect(conf event.IEventConfiguration) error {

	ctx := context.TODO()

	p.Ctx = ctx

	sub, err := conf.GetSubscriptionString()
	if err != nil {
		//Non fatal error
		log.Println(err.Error())
	} else {
		p.Subscription = sub
	}

	topic, err := conf.GetTopic()
	if err != nil {
		return nil
	}
	p.Topic = topic

	project, err := conf.GetConnectionString()
	if err != nil {
		return nil
	}
	p.Project = project

	client, err := pubsub.NewClient(ctx, p.Project)
	if err != nil {
		return err
	}

	p.ClientRef = client
	p.TopicRef = client.Topic(p.Topic)

	return nil
}

//Publish ...
func (p *PubSub) Publish(data []byte) error {

	p.TopicRef.Publish(p.Ctx, &pubsub.Message{Data: data})
	return nil
}

//Subscribe ..
func (p *PubSub) Subscribe(fn func(event.IMessage)) error {

	if p.ClientRef == nil {
		return errors.New("No client set")
	}

	sub, err := p.ClientRef.CreateSubscription(p.Ctx, p.Subscription,
		pubsub.SubscriptionConfig{Topic: p.TopicRef})

	if err != nil {
		//Do not mark as faral
		log.Println(err)
	}
	p.SubscriptionRef = sub

	cctx, cancel := context.WithCancel(p.Ctx)
	p.Cancel = cancel

	log.Println("Starting subscription...")

	sub.Receive(cctx, func(arg1 context.Context, arg2 *pubsub.Message) {

		pubSubMessage := &PubSubMessage{Message: arg2}

		fn(pubSubMessage)
	})

	return nil
}

//CancelSubscribe on any active subscription
func (p *PubSub) CancelSubscribe() {
	p.Cancel()
}
