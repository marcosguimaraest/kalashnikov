package actions

import (
	"fmt"
	"mguimara/kalashnikov/internal/trello"
	"mguimara/kalashnikov/internal/trello/objects"
)

func PostCardAction(i []string) {
	c := trello.ParseWhatsappInputToCard(i)
	res := trello.ResolveRequest(c, "cards/")
	bytes, err := trello.ResolveResponse(res)
	if err != nil {
		return
	}
	var cr *objects.CardResponse
	err = trello.ByteToObject(bytes, cr)
	if err != nil {
		return
	}
	fmt.Println("\n\n\n Card response: ", cr)
}
