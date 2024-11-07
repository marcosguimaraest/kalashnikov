package input

import (
	"go.mau.fi/whatsmeow/types/events"
)

type Commands struct {
	Command         string
	Description     string
	EveryoneAllowed bool
	Exec            func(*events.Message)
}

var ActiveCommands *[]Commands

func InitializeCommands() {
	ActiveCommands = &[]Commands{
		Commands{
			Command:         "!i - IMPRESSORA",
			Description:     "Limpeza, travada...",
			EveryoneAllowed: true,
			Exec:            nil,
		},
		Commands{
			Command:         "!p - PRODUTOS",
			Description:     "Mudar preço, cadastrar...",
			EveryoneAllowed: true,
			Exec:            nil,
		},
		Commands{
			Command:         "!c - CAIXA",
			Description:     "Caixa travado, item com erro...",
			EveryoneAllowed: true,
			Exec:            nil,
		},
		Commands{
			Command:         "!s - SISTEMA",
			Description:     "Sistema com lentidão...",
			EveryoneAllowed: true,
			Exec:            nil,
		},
	}
}
