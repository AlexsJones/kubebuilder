package processor

import (
	"encoding/base64"
	"fmt"
	"log"

	"github.com/AlexsJones/kubebuilder/src/data"
	yaml "gopkg.in/yaml.v2"
)

//NewIntentionsMapping Creates the intentions mappings...
func NewIntentionsMapping() *map[string]func(*data.Message) (bool, *data.Message) {

	return &map[string]func(*data.Message) (bool, *data.Message){
		"SYN": func(p *data.Message) (bool, *data.Message) {
			log.Printf("Receieved SYN intention \n")

			if p.Context.String() == "" {
				log.Println("Cannot SYN to message without context")
				return false, nil
			}
			reply := data.NewMessage(p.Context.String())
			reply.Type = data.Message_ACK

			return true, reply
		},
		"ACK": func(p *data.Message) (bool, *data.Message) {

			log.Printf("Receieved ACK intention with context %s\n", p.Context.Value)

			return false, nil
		},
		"STATECHANGE": func(p *data.Message) (bool, *data.Message) {
			log.Println("Receiving statechange intention")
			fmt.Println("--State Change--")
			fmt.Println(p.Payload)
			fmt.Println("----------------")
			return false, nil
		},
		"BUILD": func(p *data.Message) (bool, *data.Message) {
			log.Println("Receiving build intention")

			if p.Payload == "" {
				log.Println("Missing build payload")
				return true, data.NewStateMessage(p.Context.String(), "Missing build payload")
			}
			decoded, err := base64.StdEncoding.DecodeString(p.Payload)
			if err != nil {
				log.Println("decode error:", err)
				return true, data.NewStateMessage(p.Context.String(), "Decoding error")
			}

			builddef := data.BuildDefinition{}

			err = yaml.Unmarshal(decoded, &builddef)
			if err != nil {
				log.Fatalf("error: %v", err)
				return true, data.NewStateMessage(p.Context.String(), "Unmarshalling error")
			}
			log.Printf("%v\n", builddef)

			return false, nil
		},
	}
}
