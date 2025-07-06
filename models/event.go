package models

import (
	"fmt"
	"time"
)

type Event struct {
	ID          int
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserId      int
}

// TODO: replace with a MYSQL DB
var events = []Event{}

func (e Event) Save() {
	events = append(events, e)
	fmt.Println(events)
}

func GetAllEvents() []Event {
	return events
}
