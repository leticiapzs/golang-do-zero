package main

import "fmt"

func main() {

	// Array - tamanho fixo
	var array [2]string
	array[0] = "Hello"
	array[1] = "World"
	fmt.Println(array[0])
	fmt.Println(array[1])
	fmt.Println(array[0], array[1])
	fmt.Println(array)

	numPrimos := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(numPrimos)
	fmt.Println(numPrimos[0:4])
	fmt.Println(numPrimos[0:6])

	//Slices
	//var slice []string
	slice := make([]string, 5)
	slice[0] = "Hello"
	slice[1] = "World"
	fmt.Println(slice[0], slice[1])
	fmt.Println(slice[0])
	fmt.Println(slice[1])
	fmt.Println(slice[2])
	slice[2] = "Steph"
	fmt.Println(slice[2])
	fmt.Println(slice)

	fmt.Println(len(slice))
	fmt.Println(slice[2])
	fmt.Println(slice[3])
	fmt.Println(slice[4])
	fmt.Println(slice[5])

	numPares := []int{2, 4, 6, 8, 10, 12}
	fmt.Println(numPares)

	numPares = append(numPares, 14, 16, 18, 20)
	fmt.Println(numPares)
}

// LISTAS

// 1 - Arrays e Slices: Homogêneos
// todos os elementos tem o mesmo tipo
// [1, 2, 3, 4, 5, 6] - []int
// ["steph", "bento", "golang"] - []string

// 2 - Maps: Heterogêneos
// pode misturar tipos
// estrutura chave - valor
// [key] = value
// chave tem um tipo, e o valor pode ter outro
// map[string]int
// { "steph": 28, "bento": 4}
// map[string]string
// { "steph": "cardoso", "bento": "cardoso" }

// Array

// Tamanho fixo, de zero ou mais elementos do mesmo tipo
// Acessamos os valores com índice: a[0], a[1]...
// Função embutida len() retorna o tamanho do array
// Por conta do tamanho fixo, não é tão usado. Só em casos específicos

// Slice

// Tipo o array, mas sem tamanho fixo
// Acessamos os valores com índice: a[0], a[1]...
// Função embutida len() retorna o tamanho do slice
// Função append() usada para adicionar valores nesse slice
