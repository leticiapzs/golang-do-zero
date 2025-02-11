package main

import (
	"fmt"
)

func main() {

	// posicao := 1
	// switch posicao {
	// case 1:
	// 	fmt.Println("Primeiro lugar")
	// case 2:
	// 	fmt.Println("Segundo lugar")
	// case 3:
	// 	fmt.Println("Terceiro lugar")
	// }

	nome := "steph"
	switch nome {
	// nome == "steph"
	case "steph":
		fmt.Println("É uma pessoa")
	case "bento":
		fmt.Println("É um cachorro")
	case "bob":
		fmt.Println("É um personagem")
	}
}
