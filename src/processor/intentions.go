package processor

import (
	"encoding/base64"
	"log"

	"github.com/AlexsJones/kubebuilder/src/data"
	yaml "gopkg.in/yaml.v2"
)

//NewIntentionsMapping Creates the intentions mappings...
func NewIntentionsMapping() *map[string]func(*data.Message) (bool, *data.Message) {

	return &map[string]func(*data.Message) (bool, *data.Message){
		"SYN": func(p *data.Message) (bool, *data.Message) {
			log.Printf("Receieved ACK intention \n")

			if p.Context.String() == "" {
				log.Println("Cannot ACK to message without context")
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
			return false, nil
		},
		"BUILD": func(p *data.Message) (bool, *data.Message) {
			log.Println("Receiving build intention")

			if p.Payload == "" {
				//Could also nack malformed build here
				log.Println("Missing build payload")
				return false, nil
			}
			decoded, err := base64.StdEncoding.DecodeString(p.Payload)
			if err != nil {
				log.Println("decode error:", err)
				return false, nil
			}
			//Hand cranking a build definition for the test
			builddef := data.BuildDefinition{}

			err = yaml.Unmarshal(decoded, &builddef)
			if err != nil {
				log.Fatalf("error: %v", err)
			}
			log.Printf("--- t:\n%v\n\n", builddef)

			return false, nil
		},
	}
}
