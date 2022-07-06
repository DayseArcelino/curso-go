package main

import (
	"errors"
	"fmt"
)

func main() {
	var numero int64 = -1000000000000
	fmt.Println(numero)

	var numero2 uint32 = 10000
	fmt.Println(numero2)

	// Alias
	// Int32 = RUNE
	var numero3 rune = 12456
	fmt.Println(numero3)

	// Byte = Uint8
	var numero4 byte = 123
	fmt.Println(numero4)

	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 1230000000.45
	fmt.Println(numeroReal2)

	numeroReal3 := 12345
	fmt.Println(numeroReal3)

	// Fim números reais

	// Strings
	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	char := 'A'
	fmt.Println(char)

	// Fim strings

	texto := 5
	fmt.Println(texto)

	var boolano1 bool = true
	fmt.Println(boolano1)

	var erro error = errors.New("Erro interno")
	fmt.Println(erro)
}