package entity

import "time"

type User struct {
	ID     int
	Events map[time.Time][]*Event
}

func NewUser() *User {
	return &User{
		Events: make(map[time.Time][]*Event),
	}
}
