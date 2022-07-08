package main

import (
	"fmt"
	"modulo/Testes_Automatizados/Introducao/enderecos"
)

func main() {
	tipoEndereco := enderecos.TipoDeEndereco("Avenida Paulista")
	fmt.Println(tipoEndereco)
}
