package main

import (
	"fmt"
	"time"
)

func main() {
	go escrever("suó")
	escrever("golang")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}

}
