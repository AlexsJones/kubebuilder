package data

import (
	"fmt"
	"log"
	"time"

	"github.com/nu7hatch/gouuid"
)

//NewMessageContext returns a uuid time string for a new context
func NewMessageContext() string {
	u, err := uuid.NewV4()
	if err != nil {
		log.Fatal(err)
	}
	timeString := string(time.Now().Unix())
	return fmt.Sprintf("%s:%s", u.String(), timeString)
}

//NewMessage Creates a new bare state object
func NewMessage(context ...string) *Message {
	u, err := uuid.NewV4()
	if err != nil {
		log.Fatal(err)
	}
	senderContext := ""
	if len(context) > 0 {
		senderContext = context[0]
	}
	return &Message{Uuid: &UUID{Value: u.String()}, Context: &UUID{Value: senderContext}}

}
