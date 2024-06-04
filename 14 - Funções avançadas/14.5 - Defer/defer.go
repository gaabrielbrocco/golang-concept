package main

import "fmt"

func funcao1() {
	fmt.Println("func 1")
}

func funcao2() {
	fmt.Println("func 2")
}

func alunoAprovado(n1, n2 float32) bool {
	fmt.Println("funcao pra verificar aluno")

	media := (n1 + n2) / 2

	if media >= 6 {
		return true
	}
	return false
}

func main() {
	defer funcao1()
	funcao2()

	fmt.Println(alunoAprovado(7, 8))
}
