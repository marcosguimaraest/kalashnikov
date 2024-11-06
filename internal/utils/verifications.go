package utils

import (
	"fmt"

	"go.mau.fi/whatsmeow/types"
)

func VerifyNumber(mi types.MessageInfo) bool {
	if mi.Sender.User == "5521970079469" {
		return (true)
	}
	fmt.Println("\n\n\n wrong nbr")
	return (false)
}
