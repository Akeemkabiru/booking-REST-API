package models

import "time"

type Event struct{
	ID int    
	Name string   `binding:"required"`
	Description string `binding:"required"`
	Location string `binding:"required"`
	DateTime time.Time 
	UserId int
}

var event []Event
func (e Event) Save() []Event{
   event = append(event, e)
   return event
}

func GetAllEvents() []Event{
  return event
}

