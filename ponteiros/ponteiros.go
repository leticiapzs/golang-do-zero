package main

import "fmt"

// Ponteiros: Um ponteiro nada mais é do que uma variável que ao invés
// de armazenar um valor (1, "olá", true..), armazena um end de memória

func zeroValue(i int) {
	i = 0
	fmt.Println("End de memória do i dentro da func: ", &i)
}

func zeroPointer(i *int) {
	*i = 0
}

func main() {
	// Mémoria -> essa memória tem um endereço -> esse endereço guarda um valor
	i := 1
	fmt.Println("Valor inicial:", i)
	fmt.Println("Valor end de memória: ", &i) //&var -> apontando para o end de memória da var

	zeroValue(i)
	fmt.Println("zeroval:", i)

	zeroPointer(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)

	// a := &i
	// fmt.Println("Variavel a:", a)
	// fmt.Println("Variavel *a:", *a) // quando uso *, estamos falando do endereço! E não do valor
	// fmt.Println("Variavel &a:", &a)
}
