package handlers

import (
	"mguimara/kalashnikov/internal/helpers"
	"mguimara/kalashnikov/internal/input"
	"strings"

	"go.mau.fi/whatsmeow/types/events"
)

func MessageHandler(m *events.Message) {
	messageText := m.Message.GetConversation()
	splitedMessage := strings.Split(messageText, " ")
	var replyMessage string
	for _, c := range *input.ActiveCommands {
		if splitedMessage[0] == c.Input {
			r := input.Restaurants[splitedMessage[1]]
			if r == "" {
				helpers.Reply(m, "Restaurante inválido")
				return
			}
			replyMessage += "RESTAURANTE: " + r + "\n"
			replyMessage += "DESCRIÇÃO: " + strings.Join(splitedMessage[2:], " ")
			helpers.Reply(m, replyMessage)
			return
		}
	}
	helpers.Reply(m, input.GetMenu())
}
