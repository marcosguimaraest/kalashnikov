package input

import (
	"go.mau.fi/whatsmeow/types/events"
)

type Commands struct {
	Input           string
	Command         string
	Description     string
	EveryoneAllowed bool
	Exec            func(*events.Message)
}

var ActiveCommands *[]Commands

func InitializeCommands() {
	ActiveCommands = &[]Commands{
		Commands{
			Input:           "!i",
			Command:         "!i - IMPRESSORA",
			Description:     "Limpeza, travada...",
			EveryoneAllowed: true,
			Exec:            nil,
		},
		Commands{
			Input:           "!p",
			Command:         "!p - PRODUTOS",
			Description:     "Mudar preço, cadastrar...",
			EveryoneAllowed: true,
			Exec:            nil,
		},
		Commands{
			Input:           "!c",
			Command:         "!c - CAIXA",
			Description:     "Caixa travado, item com erro...",
			EveryoneAllowed: true,
			Exec:            nil,
		},
		Commands{
			Input:           "!s",
			Command:         "!s - SISTEMA",
			Description:     "Sistema com lentidão...",
			EveryoneAllowed: true,
			Exec:            nil,
		},
	}
}
