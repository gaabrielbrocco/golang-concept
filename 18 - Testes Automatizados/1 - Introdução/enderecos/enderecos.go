package enderecos

import "strings"

// TipoDeEndereco verifica se o endereço tem tipo válido
func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"Rua", "Avenida", "Estrada", "Rodovia"}
	letraMinuscula := strings.ToLower(endereco)
	primeiraPalavra := strings.Split(letraMinuscula, " ")[0]

	enderecoTemUmTipoValido := false
	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavra {
			enderecoTemUmTipoValido = true
		}
	}

	if enderecoTemUmTipoValido {
		return strings.Title(primeiraPalavra)
	}

	return "Tipo inválido"
}
