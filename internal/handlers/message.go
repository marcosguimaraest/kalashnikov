package handlers

import (
	"mguimara/kalashnikov/internal/helpers"
	"mguimara/kalashnikov/internal/input"
	"mguimara/kalashnikov/internal/trello/actions"
	"strings"

	"go.mau.fi/whatsmeow/types/events"
)

func MessageHandler(m *events.Message) {
	messageText := m.Message.GetConversation()
	splittedMessage := strings.Split(messageText, " ")
	var replyMessage string
	for _, c := range *input.ActiveCommands {
		if splittedMessage[0] == c.Input {
			r := input.Restaurants[splittedMessage[1]]
			if r == "" && splittedMessage[2] == "" {
				helpers.Reply(m, "Restaurante inválido")
				return
			}
			replyMessage += "RESTAURANTE: " + r + "\n"
			replyMessage += "DESCRIÇÃO: " + strings.Join(splittedMessage[2:], " ")
			actions.PostCardAction(splittedMessage[2:])
			helpers.Reply(m, replyMessage)
			return
		}
	}
	helpers.Reply(m, input.GetMenu())
}
