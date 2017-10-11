package processor

import (
	"log"

	event "github.com/AlexsJones/cloud-transponder/events"
	"github.com/AlexsJones/kubebuilder/src/data"
	"github.com/golang/protobuf/proto"
)

//MessageProcessor object definition
type MessageProcessor struct {
	intentionMap           *map[string]func(p *data.Message)
	PubSubRef              event.IEvent
	PubSubConfigurationRef event.IEventConfiguration
}

//NewMessageProcessor creates a new MessageProcessor object
func NewMessageProcessor(intentions *map[string]func(p *data.Message), pubsub event.IEvent, pubsubconf event.IEventConfiguration) *MessageProcessor {
	return &MessageProcessor{
		intentionMap:           intentions,
		PubSubRef:              pubsub,
		PubSubConfigurationRef: pubsubconf,
	}
}

//Start the blocking process to ingest and egest traffic
func (m *MessageProcessor) Start() {
	if err := event.Subscribe(m.PubSubRef, func(arg2 event.IMessage) {

		if ok, err := m.ingest(arg2); err != nil {
			//Currently no handler for a failed message
		} else {
			if ok {
				arg2.Ack()
			} else {
				arg2.Nack()
			}
		}
	}); err != nil {
		log.Fatal(err)
	}
}

//Ingest and processes incoming messages
func (m *MessageProcessor) ingest(message event.IMessage) (bool, error) {
	log.Printf("Received message of size %d\n", len(message.GetRaw()))
	st := &data.Message{}
	if err := proto.Unmarshal(message.GetRaw(), st); err != nil {
		log.Fatalln("Failed to parse message...", err)
		return false, err
	}

	im := m.intentionMap

	if fn, ok := (*im)[st.Type.String()]; ok {
		fn(st)
	}
	log.Println(st.Type.String())

	return true, nil
}

func (m *MessageProcessor) egest(message event.IMessage) (bool, error) {

	return false, nil
}
