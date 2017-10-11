package processor

import (
	"log"

	event "github.com/AlexsJones/cloud-transponder/events"
	"github.com/AlexsJones/kubebuilder/src/data"
	"github.com/golang/protobuf/proto"
)

//MessageProcessor processes incoming messages
func MessageProcessor(message event.IMessage) (bool, error) {
	log.Println("Received message...")
	st := &data.Message{}
	if err := proto.Unmarshal(message.GetRaw(), st); err != nil {
		log.Fatalln("Failed to parse message...", err)
		return false, err
	}

	log.Println(st.GetUuid().String())

	return true, nil
}
