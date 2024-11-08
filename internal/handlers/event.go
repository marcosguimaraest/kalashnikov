package handlers

import (
	"encoding/json"
	"fmt"
	"mguimara/kalashnikov/internal/appdb/services"
	"mguimara/kalashnikov/internal/models"

	"go.mau.fi/whatsmeow/types/events"
)

func EventHandler(evt interface{}) {
	switch v := evt.(type) {
	case *events.Message:
		//Check if user is associated
		marshal, err := json.Marshal(v.Message)
		if err == nil {
			fmt.Println("\n\n\nMESSAGE\n", string(marshal), "\n\n\n")
		} else {
			fmt.Println("\n\nERROR\n\n")
		}
		services.InsertUser(models.User{NumberID: "5521991524379", Name: "MARCOS GUIMARAES"})
		_, err = services.GetUser(v.Info.Sender.User)
		if err == nil && !v.Info.IsGroup && v.Info.Chat.User == "5521991524379" {
			MessageHandler(v)
		}
	}
}
