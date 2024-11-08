package handlers

import (
	"mguimara/kalashnikov/internal/appdb/services"

	"go.mau.fi/whatsmeow/types/events"
)

func EventHandler(evt interface{}) {
	switch v := evt.(type) {
	case *events.Message:
		//Check if user is associated
		_, err := services.GetUser(v.Info.Sender.User)
		if err == nil && !v.Info.IsGroup && v.Info.Chat.User == "5521991524379" {
			MessageHandler(v)
		}
	}
}
