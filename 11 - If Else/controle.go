package main

import "fmt"

func main() {
	fmt.Println()

	numero := -2

	if numero > 15 {
		fmt.Println("maior q 15")
	} else {
		fmt.Println("menor q 15")
	}

	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("é maior q zero")
	} else {
		fmt.Println("é menor q zero")
	}

}
