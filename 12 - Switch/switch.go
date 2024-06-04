package main

import "fmt"

func diaSemana(numero int) string {
	switch numero {
	case 1:
		return "domingo"
	case 2:
		return "segunda"
	case 3:
		return "terça"
	case 4:
		return "quarta"
	case 5:
		return "quinta"
	case 6:
		return "sexta"
	case 7:
		return "sabado"
	default:
		return "numero inválido"
	}

}

func diaSemana2(numero int) string {
	switch {
	case numero == 1:
		return "domingo"
	default:
		return "invalido"
	}
}

func main() {
	fmt.Println()

	dia := diaSemana(3)
	fmt.Println(dia)

}
