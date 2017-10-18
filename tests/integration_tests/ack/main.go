package main

import (
	"log"
	"time"

	event "github.com/AlexsJones/cloud-transponder/events"
	"github.com/AlexsJones/cloud-transponder/events/pubsub"
	"github.com/AlexsJones/kubebuilder/src/config"
	"github.com/AlexsJones/kubebuilder/src/data"
	"github.com/golang/protobuf/proto"
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

	gconfig.Topic = conf.PubSubConfiguration.Topic
	gconfig.ConnectionString = conf.PubSubConfiguration.ConnectionString

	if err = event.Connect(gpubsub, gconfig); err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 10; i++ {
		//Generate a new state object
		st := data.NewMessage(data.NewMessageContext())

		out, err := proto.Marshal(st)
		if err != nil {
			log.Fatalln("Failed to encode address book:", err)
		}
		event.Publish(gpubsub, out)
	}

	time.Sleep(time.Minute)
}
