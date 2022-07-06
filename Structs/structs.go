package main

import (
	"fmt"
)

type usuario struct {
	nome string
	idade uint8
}

func main() {
	fmt.Println("Arquivo Structs")

	var u usuario
	u.nome = "Pedro"
	u.idade = 21
	fmt.Println(u)

	/* Ou pode ser feito da seguinte maneira:

	usuario2 := usuario{"Pedro", 21}
	fmt.Println(usuario2)
	*/ 


}