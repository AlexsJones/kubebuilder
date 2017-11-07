package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	event "github.com/AlexsJones/cloud-transponder/events"
	"github.com/AlexsJones/cloud-transponder/events/pubsub"
	"github.com/AlexsJones/kubebuilder/src/config"
	"github.com/AlexsJones/kubebuilder/src/log"
	"github.com/AlexsJones/kubebuilder/src/processor"
	validator "gopkg.in/go-playground/validator.v9"
)

func main() {

	confPath := flag.String("conf", "", "Set a custom configuration path")

	flag.Parse()

	path := "test-config.yaml"
	if strings.Compare(*confPath, "") != 0 {
		path = *confPath
		logger.GetInstance().Log(fmt.Sprintf("Using custom config path %s", path))
	} else {
		logger.GetInstance().Log("Using test-config.yaml")
	}
	//Load configuration
	conf, err := config.LoadConfiguration(path)
	if err != nil {
		log.Fatal(err)
	}
	//Validate configuration structure

	validate := validator.New()
	err = validate.Struct(conf)
	if err != nil {

		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like this.
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return
		}
	}
	//Create our GCP pubsub
	gpubsub := gcloud.NewPubSub()

	//Create the GCP Pubsub configuration
	gconfig := gcloud.NewPubSubConfiguration()

	gconfig.Topic = conf.PubSubConfiguration.Topic
	gconfig.ConnectionString = conf.PubSubConfiguration.ConnectionString
	gconfig.SubscriptionString = conf.PubSubConfiguration.SubscriptionString

	if err := event.Connect(gpubsub, gconfig); err != nil {
		logger.GetInstance().Fatal(err.Error())
	}
	//Create a message processor
	messageProcessor := processor.NewMessageProcessor(processor.NewIntentionsMapping(), gpubsub, gconfig, conf)

	//Check our debug mode for queue draining
	if conf.KubeBuilderConfiguration.Drainqueue {
		logger.GetInstance().Log("Draining queue")
		messageProcessor.Drain()
	} else {
		messageProcessor.Start()
	}
}
