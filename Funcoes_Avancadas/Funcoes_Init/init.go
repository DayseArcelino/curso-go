package main

import "fmt"

func init() {
	fmt.Println("Executando a função init")
}

func main() {
	fmt.Println("Função main sendo executada")
}

// Não importa se a função init vim antes ou depois da função main, ela sempre será executada primeiro.