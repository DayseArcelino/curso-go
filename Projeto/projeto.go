package main

import "fmt"

type usuario struct {
	nome string
	nota uint8
	nota1 uint8
	nota2 uint8
}

func alunosEstaAprovado(n1, n2, n3 float32) bool {

	media := (n1 + n2 + n3) / 3

	if media >= 7 {
		return true
	}
	return false
}

/* func nome(txt string) string {

	nome := usuario("João", "Maria", "José")

	return nome
} */

func main() {
	João := usuario{"João", 10, 8, 5}

	Maria := usuario{"Maria", 7, 5, 8}

	José := usuario{"José", 4, 0, 8}

	aprovado := usuario.Aprovado()
	fmt.Println(aprovado)

	fmt.Println(alunosEstaAprovado(7, 8, 6))
}
