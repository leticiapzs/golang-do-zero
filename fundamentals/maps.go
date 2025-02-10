package main

import "fmt"

func main() {
	// idade := map[string]int{}
	// idade["steph"] = 28
	// idade["bento"] = 4
	// fmt.Println(idade)
	// fmt.Println(idade["steph"])
	// fmt.Println(idade["bento"])

	anoNasc := map[string]int{
		"steph": 1995,
		"bento": 2008,
	}
	fmt.Println(anoNasc)
	fmt.Println(anoNasc["steph"])
	fmt.Println(anoNasc["bento"])
	anoNasc["golangDoZero"] = 2024
	fmt.Println(anoNasc)
}

// 2 - Maps: Heterogêneos
// pode misturar tipos
// estrutura chave - valor
// [key] = value
// chave tem um tipo, e o valor pode ter outro
// map[k]v -> k = chave, v = valor

// map[string]int
// { "steph": 28, "bento": 4}
// map[string]string
// { "steph": "cardoso", "bento": "cardoso" }
