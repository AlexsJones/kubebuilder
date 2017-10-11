package processor

import (
	"log"

	"github.com/AlexsJones/kubebuilder/src/data"
)

//NewIntentionsMapping Creates the intentions mappings...
func NewIntentionsMapping() *map[string]func(*data.Message) {

	return &map[string]func(*data.Message){
		"SYN": func(p *data.Message) {
			log.Println("Received SYN message...")
		},
		"ACK": func(p *data.Message) {

		},
	}
}
