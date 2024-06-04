package main

import "fmt"

func closure() func() {
	texto := "dentro da closure"

	funcao := func() {
		fmt.Println((texto))
	}

	return funcao
}

func main() {
	texto := "dentro da main"
	fmt.Println(texto)

	funcaoNova := closure()
	funcaoNova()
}
