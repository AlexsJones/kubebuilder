package data

import (
	"log"

	"github.com/nu7hatch/gouuid"
)

//NewMessage Creates a new bare state object
func NewMessage() *Message {
	u, err := uuid.NewV4()
	if err != nil {
		log.Fatal(err)
	}
	return &Message{Uuid: &UUID{Value: u.String()}}

}
