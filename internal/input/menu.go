package input

func GetMenu() string {
	menu := "Obrigado por entrar em contato, \n\n *SELECIONE*\n\n"
	for _, c := range *ActiveCommands {
		menu += "*" + c.Command + "* - " + c.Description + "\n"
	}
	return (menu)
}