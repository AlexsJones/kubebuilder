package data

import (
	"log"

	"github.com/nu7hatch/gouuid"
)

//NewState Creates a new bare state object
func NewState() *State {
	u, err := uuid.NewV4()
	if err != nil {
		log.Fatal(err)
	}
	return &State{Uuid: &UUID{Value: u.String()}}

}
