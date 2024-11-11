package objects

import "encoding/json"

type Card struct {
	Name   string `json:"name"`
	Desc   string `json:"desc"`
	IDList string `json:"idList"`
}

func ByteToCard(b []byte) *Card {
	var c Card
	err := json.Unmarshal(b, &c)
	if err != nil {
		return nil
	}
	return (&c)
}
