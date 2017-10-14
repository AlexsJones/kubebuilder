package processor

import (
	"log"
	"time"

	event "github.com/AlexsJones/cloud-transponder/events"
	"github.com/AlexsJones/kubebuilder/src/data"
	"github.com/AlexsJones/kubebuilder/src/fabricarium"
	"github.com/golang/protobuf/proto"
)

//MessageProcessor object definition
type MessageProcessor struct {
	intentionMap           *map[string]func(p *data.Message) (bool, *data.Message)
	PubSubRef              event.IEvent
	PubSubConfigurationRef event.IEventConfiguration
	FabricariumRef         *fabricarium.Fabricarium
}

//NewMessageProcessor creates a new MessageProcessor object
func NewMessageProcessor(intentions *map[string]func(p *data.Message) (bool, *data.Message), pubsub event.IEvent, pubsubconf event.IEventConfiguration,
	fabricarium *fabricarium.Fabricarium) *MessageProcessor {
	return &MessageProcessor{
		intentionMap:           intentions,
		PubSubRef:              pubsub,
		PubSubConfigurationRef: pubsubconf,
	}
}

//Start the blocking process to ingest and egest traffic
func (m *MessageProcessor) Start() {

	messageChannel := make(chan *data.Message)

	go func() {
		if err := event.Subscribe(m.PubSubRef, func(arg2 event.IMessage) {

			if ok, err := m.ingest(arg2, messageChannel); err != nil {
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
	}()

	for {
		select {
		case msg := <-messageChannel:
			if msg != nil {
				log.Println("Message processed & replied")
				m.egest(msg)
			} else {
				log.Println("Message processed")
			}
		default:
		}
		time.Sleep(time.Millisecond * 500)
	}

}

//Ingest and processes incoming messages
func (m *MessageProcessor) ingest(message event.IMessage, messageChannel chan *data.Message) (bool, error) {
	log.Printf("Received message of size %d\n", len(message.GetRaw()))
	st := &data.Message{}
	if err := proto.Unmarshal(message.GetRaw(), st); err != nil {
		log.Fatalln("Failed to parse message...", err)
		return false, err
	}

	im := m.intentionMap

	if fn, ok := (*im)[st.Type.String()]; ok {

		if reply, parcel := fn(st); reply {
			messageChannel <- parcel
		} else {
			messageChannel <- nil
		}
	}
	return true, nil
}

//egest messages outward into the pubsub
func (m *MessageProcessor) egest(message proto.Message) (bool, error) {

	out, err := proto.Marshal(message)
	if err != nil {
		log.Printf("Failed to encode message %s\n", err.Error())
		return false, err
	}

	if err := event.Publish(m.PubSubRef, out); err != nil {
		return false, err
	}

	return true, nil
}
