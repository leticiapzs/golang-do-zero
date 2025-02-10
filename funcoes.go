package main

import "fmt"

func main() {
	somaDosValores := soma(42, 13)
	fmt.Println(somaDosValores)

	nome, sobrenome := printaNomeCompleto("STEPH", "CARDOSO")
	fmt.Println(nome)
	fmt.Println(sobrenome)
}

// Função começando com letra minúscula:
// Função é PRIVADA
// Ela só pode ser utilizada no próprio pacote
func printaNomeCompleto(nome, sobrenome string) (string, string) {
	return nome, sobrenome
}

// Função começando com letra maiuscula:
// Função é PÚBLICA
// Ela pode ser utilizada fora do próprio pacote
// Como utilizaria ela fora do pacote: main.PrintaNome(nome)
func PrintaNome(nome string) string {
	return nome
}

func soma(x int, y int) int {
	return x + y
}
