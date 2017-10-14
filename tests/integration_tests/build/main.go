package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	event "github.com/AlexsJones/cloud-transponder/events"
	"github.com/AlexsJones/cloud-transponder/events/pubsub"
	"github.com/AlexsJones/kubebuilder/src/config"
	"github.com/AlexsJones/kubebuilder/src/data"
	"github.com/golang/protobuf/proto"
	yaml "gopkg.in/yaml.v2"
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

	if err = event.Connect(gpubsub, gconfig); err != nil {
		log.Fatal(err)
	}

	//Generate a new state object
	st := data.NewMessage(data.NewMessageContext())
	//Set our outbound message to indicate a build
	st.Type = data.Message_BUILD

	//Load yaml
	raw, err := ioutil.ReadFile("./tests/integration_tests/build/testbuild.yaml")
	if err != nil {
		log.Fatal(err)
	}

	//Hand cranking a build definition for the test
	builddef := data.BuildDefinition{}

	err = yaml.Unmarshal(raw, &builddef)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- t:\n%v\n\n", builddef)

	//Add the build as an encoded string into our message
	out, err := yaml.Marshal(builddef)
	if err != nil {
		log.Fatalln("Failed to marshal:", err)
	}

	st.Payload = base64.StdEncoding.EncodeToString(out)

	out, err = proto.Marshal(st)
	if err != nil {
		log.Fatalln("Failed to encode:", err)
	}

	event.Publish(gpubsub, out)

	time.Sleep(time.Minute)
}
