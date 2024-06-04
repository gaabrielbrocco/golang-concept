package main

import "fmt"

func main() {
	canal := make(chan string, 2)

	canal <- "hello"

	mensagem := <-canal
	fmt.Println(mensagem)
}
