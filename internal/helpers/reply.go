package helpers

import (
	"context"
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
		/*
			ListMessage: &wp.ListMessage{
				Title:       proto.String(text),
				ListType:    wp.ListMessage_SINGLE_SELECT.Enum(),
				ContextInfo: &wp.ContextInfo{},
				Sections: []*wp.ListMessage_Section{
					&wp.ListMessage_Section{
						Rows: []*wp.ListMessage_Row{
							&wp.ListMessage_Row{
								Title:       proto.String("MENU 1"),
								Description: proto.String("MENU 1"),

								RowID: proto.String("1"),
							}, &wp.ListMessage_Row{
								Title:       proto.String("MENU 2"),
								Description: proto.String("MENU 2"),
								RowID:       proto.String("2"),
							},
							&wp.ListMessage_Row{
								Title:       proto.String("MENU 3"),
								Description: proto.String("MENU 3"),
								RowID:       proto.String("3"),
							},
						},
					},
				},
			},*/
	}

	_, err := client.KalashnikovClient.SendMessage(context.Background(), msg.Info.Chat, newMessage)
	if err != nil {
		fmt.Println("\n\n\n ERROR\n", err)
	}
}
