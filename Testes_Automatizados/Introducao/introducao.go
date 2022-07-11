package main

import (
	"fmt"
	"modulo/Testes_Automatizados/Introducao/enderecos"
)

func main() {
	tipoEndereco := enderecos.TipoDeEndereco("Rodovia dos Imigrantes")
	fmt.Println(tipoEndereco)
}
