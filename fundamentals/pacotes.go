package main

import (
	"fmt"
	steph "strings"
)

func main() {
	fmt.Println("Hello, world!")

	// sem alterar o nome do pacote,
	// usaríamos chamando strings.Func()
	// strings.Slipt()
	fmt.Println(steph.Split("steph", ""))
}
