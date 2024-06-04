package main

import "fmt"

func main() {

	retorno := func(texto string) string {
		return fmt.Sprintf("recebido -> %s", texto)
	}("Passando parametro")

	fmt.Println(retorno)
}
