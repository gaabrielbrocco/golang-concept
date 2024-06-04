package main

import (
	"fmt"
	"time"
)

func main() {
	// i := 0

	// for i < 10 {
	// 	i++
	// 	fmt.Println("incrementando i")
	// 	time.Sleep(time.Second)
	// }

	// fmt.Println(i)
	fmt.Println("-------------")

	for j := 0; j < 10; j++ {
		fmt.Println("incrementando j", j)
		time.Sleep(time.Second)
	}

	nomes := [3]string{"joao", "davi", "lucas"}

	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}
}
