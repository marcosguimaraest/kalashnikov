package input

func GetMenu() string {
	menu := "Obrigado por entrar em contato, \n\n *SELECIONE*\n\n"
	for _, c := range *ActiveCommands {
		menu += "*" + c.Command + "* - " + c.Description + "\n"
	}
	menu += "\n *EXEMPLOS*\n"
	menu += "\n _!i ROASTERY A impressora do bar está com intermitência!_\n"
	menu += "\n _!s SQUARE Sistema muito lento... restaurante lotado! ueueunneuenn (choro)_"
	return (menu)
}
