package helpers

import (
	"context"
	"fmt"
	"mguimara/kalashnikov/internal/client"

	wp "go.mau.fi/whatsmeow/proto/waE2E"
	"go.mau.fi/whatsmeow/types/events"
	proto "google.golang.org/protobuf/proto"
)

func Reply(msg *events.Message) {
	var newMessage = &wp.Message{
		ExtendedTextMessage: &wp.ExtendedTextMessage{
			Text: proto.String("Teste"),
		},
	}

	_, err := client.KalashnikovClient.SendMessage(context.Background(), msg.Info.Chat, newMessage)
	if err != nil {
		fmt.Println("\n\n\n ERROR\n")
	}
}
