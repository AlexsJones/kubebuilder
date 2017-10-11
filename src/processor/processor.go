package processor

import (
	"log"

	event "github.com/AlexsJones/cloud-transponder/events"
	"github.com/AlexsJones/kubebuilder/src/data"
	"github.com/golang/protobuf/proto"
)

//MessageProcessor object definition
type MessageProcessor struct {
	intentionMap *map[string]func(p *data.Message)
}

//NewMessageProcessor creates a new MessageProcessor object
func NewMessageProcessor(intentions *map[string]func(p *data.Message)) *MessageProcessor {
	return &MessageProcessor{
		intentionMap: intentions,
	}
}

//Ingest and processes incoming messages
func (m *MessageProcessor) Ingest(message event.IMessage) (bool, error) {
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
