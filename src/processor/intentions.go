package processor

import (
	"encoding/base64"
	"fmt"
	"log"

	"github.com/AlexsJones/kubebuilder/src/config"
	"github.com/AlexsJones/kubebuilder/src/data"
	"github.com/AlexsJones/kubebuilder/src/fabricarium"
	"github.com/AlexsJones/kubebuilder/src/log"
	yaml "gopkg.in/yaml.v2"
)

//NewIntentionsMapping Creates the intentions mappings...
func NewIntentionsMapping() *map[string]func(*config.Configuration, *data.Message) {

	return &map[string]func(*config.Configuration, *data.Message){
		"SYN": func(appConfig *config.Configuration, p *data.Message) {
			logger.GetInstance().Log("Receieved SYN intention")

			if p.Context.String() == "" {
				logger.GetInstance().Log("Cannot SYN to message without context")
				return
			}
			reply := data.NewMessage(p.Context.String())
			reply.Type = data.Message_ACK

		},
		"ACK": func(appConfig *config.Configuration, p *data.Message) {

			logger.GetInstance().Log(fmt.Sprintf("Receieved ACK intention with context %s", p.Context.Value))

		},
		"STATECHANGE": func(appConfig *config.Configuration, p *data.Message) {
			logger.GetInstance().Log("Receiving statechange intention")
			logger.GetInstance().Log("--State Change--")
			logger.GetInstance().Log(fmt.Sprintf("%v", p.Payload))
			logger.GetInstance().Log("----------------")
		},
		"BUILD": func(appConfig *config.Configuration, p *data.Message) {
			logger.GetInstance().Log("Receiving build intention")

			if p.Payload == "" {
				logger.GetInstance().Log("Missing build payload")

			}
			decoded, err := base64.StdEncoding.DecodeString(p.Payload)
			if err != nil {
				log.Println("decode error:", err)

			}
			builddef := data.BuildDefinition{}

			err = yaml.Unmarshal(decoded, &builddef)
			if err != nil {
				logger.GetInstance().Fatal(fmt.Sprintf("error: %v", err))
			}
			logger.GetInstance().Log(fmt.Sprintf("%v", builddef))

			fab := fabricarium.NewFabricarium(&fabricarium.Configuration{MountInformation: &fabricarium.Mount{Path: appConfig.KubeBuilderConfiguration.Mount}, ApplicationConfiguration: appConfig})

			fab.Process(&builddef)
		},
	}
}
