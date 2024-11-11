package trello

import (
	"mguimara/kalashnikov/internal/trello/api"
	"mguimara/kalashnikov/internal/trello/objects"
)

func ParseWhatsappInputToCard(a []string) objects.Card {
	var c objects.Card

	c.Name = a[0]
	for _, v := range a {
		c.Desc += v
	}
	c.IDList = api.ApiIDDefaultList
	return (c)
}
