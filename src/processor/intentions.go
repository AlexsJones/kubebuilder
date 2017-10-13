package processor

import (
	"log"

	"github.com/AlexsJones/kubebuilder/src/data"
)

//NewIntentionsMapping Creates the intentions mappings...
func NewIntentionsMapping() *map[string]func(*data.Message) (bool, *data.Message) {

	return &map[string]func(*data.Message) (bool, *data.Message){
		"SYN": func(p *data.Message) (bool, *data.Message) {
			log.Println("Received SYN message...")

			if p.Context.String() == "" {
				log.Println("Cannot ACK to message without context")
				return false, nil
			}
			reply := data.NewMessage(p.Context.String())
			reply.Type = data.Message_ACK

			return true, reply
		},
		"ACK": func(p *data.Message) (bool, *data.Message) {

			log.Printf("Receieved ACK from builder with context %s\n", p.Context.Value)
			return false, nil
		},
		"STATECHANGE": func(p *data.Message) (bool, *data.Message) {

			return false, nil
		},
		"BUILD": func(p *data.Message) (bool, *data.Message) {

			return false, nil
		},
	}
}
