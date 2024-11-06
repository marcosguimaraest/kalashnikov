package models

import (
	"go.mau.fi/whatsmeow/types"
)

type User struct {
	NumberID string `json:"numberID"`
	Name     string `json:"name"`
	Ticket   Ticket `json:"ticket"`
}

//Creates a user with message info

func NewUser(mi types.MessageInfo) User {
	return User{
		NumberID: mi.Sender.User,
		Name:     mi.PushName,
	}
}
