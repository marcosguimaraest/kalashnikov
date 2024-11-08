package helpers

import (
	"context"
	"encoding/json"
	"fmt"
	"mguimara/kalashnikov/internal/client"

	wp "go.mau.fi/whatsmeow/proto/waE2E"
	"go.mau.fi/whatsmeow/types/events"
	proto "google.golang.org/protobuf/proto"
)

func Reply(msg *events.Message, text string) {
	var newMessage = &wp.Message{
		ExtendedTextMessage: &wp.ExtendedTextMessage{
			Text: proto.String(text),
		},
	}
	marshal, err := json.Marshal(newMessage)
	if err == nil {
		fmt.Println("\n\n\nMESSAGE\n", string(marshal), "\n\n\n")
	} else {
		fmt.Println("\n\nERROR\n\n")
	}
	_, err = client.KalashnikovClient.SendMessage(context.Background(), msg.Info.Chat, newMessage)
	if err != nil {
		fmt.Println("\n\n\n ERROR\n", err)
	}
}
