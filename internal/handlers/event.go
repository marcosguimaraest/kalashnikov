package handlers

import (
	"fmt"
	"mguimara/kalashnikov/internal/appdb/services"
	"mguimara/kalashnikov/internal/helpers"
	"mguimara/kalashnikov/internal/input"
	"mguimara/kalashnikov/internal/models"

	"go.mau.fi/whatsmeow/types/events"
)

func EventHandler(evt interface{}) {
	switch v := evt.(type) {
	case *events.Message:
		err := services.InsertUser(models.User{NumberID: "5521991524379", Name: "MARCOS GUIMARAES TRINDADE"})
		if err != nil {
			fmt.Println("\n\n\nERROR - ", err)
		}
		u, err := services.GetUser(v.Info.Sender.User)
		if err == nil && !v.Info.IsGroup && v.Info.Chat.User == "5521991524379" {
			helpers.Reply(v, input.GetMenu())
			fmt.Println("\n\n\nUser - ", u)
		}
	}
}
