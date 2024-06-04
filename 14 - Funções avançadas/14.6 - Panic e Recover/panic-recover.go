package main

import "fmt"

func recuperaExecucao() {
	if r := recover(); r != nil {
		fmt.Println("recuperar execucao")
	}

}
func alunoAprovado(n1, n2 float64) bool {
	defer recuperaExecucao()
	media := (n1 + n2) / 2

	if media < 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("a media é exatamente 6")
}

func main() {
	fmt.Println(alunoAprovado(6, 6))
	fmt.Println("pós execução")
}
