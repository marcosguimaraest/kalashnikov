package models

type Ticket struct {
	ID     string `json:"id"`
	IDList string `json:"idList"`
	Desc   string `json:"desc"`
	Name   string `json:"name"`
}
