package processor

import (
	"fmt"
	"log"

	event "github.com/AlexsJones/cloud-transponder/events"
	"github.com/AlexsJones/kubebuilder/src/data"
	"github.com/AlexsJones/kubebuilder/src/log"
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

func (m *MessageProcessor) Drain() {
	for {
		event.Subscribe(m.PubSubRef, func(arg2 event.IMessage) {
			logger.GetInstance().Log("Acking message in queue drain")
			arg2.Ack()
		})
	}
}

//Ingest and processes incoming messages
func (m *MessageProcessor) ingest(message event.IMessage) (bool, error) {
	logger.GetInstance().Log(fmt.Sprintf("Received message of size %d", len(message.GetRaw())))
	st := &data.Message{}
	if err := proto.Unmarshal(message.GetRaw(), st); err != nil {
		logger.GetInstance().Fatal("Failed to parse message...")
		return false, err
	}

	im := m.intentionMap

	if fn, ok := (*im)[st.Type.String()]; ok {

		fn(st)
	}
	return true, nil
}

//egest messages outward into the pubsub
func (m *MessageProcessor) egest(message proto.Message) (bool, error) {

	out, err := proto.Marshal(message)
	if err != nil {
		logger.GetInstance().Log(fmt.Sprintf("Failed to encode message %s", err.Error()))
		return false, err
	}

	if err := event.Publish(m.PubSubRef, out); err != nil {
		return false, err
	}

	return true, nil
}
