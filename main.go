package main

import (
	"log"

	event "github.com/AlexsJones/cloud-transponder/events"
	"github.com/AlexsJones/cloud-transponder/events/pubsub"
	"github.com/AlexsJones/kubebuilder/src/config"
)

func main() {

	//Load configuration
	conf, err := config.LoadConfiguration("config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	//Create our GCP pubsub
	gpubsub := gcloud.NewPubSub()

	//Create the GCP Pubsub configuration
	gconfig := gcloud.NewPubSubConfiguration()

	gconfig.Topic = conf.GCPConfiguration.Topic
	gconfig.ConnectionString = conf.GCPConfiguration.ConnectionString
	gconfig.SubscriptionString = conf.GCPConfiguration.SubscriptionString

	if err := event.Connect(gpubsub, gconfig); err != nil {
		log.Fatal(err)
	}

	if err := event.Subscribe(gpubsub, func(arg2 event.IMessage) {

		log.Printf("Received: %s", string(arg2.GetRaw()))
		arg2.Ack()
	}); err != nil {
		log.Fatal(err)
	}
}