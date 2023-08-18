package main

import (
	"errors"
	"fmt"
)

func main() {
	var numero int16 = 100

	fmt.Println(numero)

	// rune = uint32

	var numero2 uint32 = 10000
	fmt.Println(numero2)


	var numero3 byte = 100
	fmt.Println(numero3)


	var nR1 float32 = 12334444.5
	fmt.Println(nR1)

	var nR2 float64 = 11111111111111111.3
	fmt.Println(nR2)


	str2 := "Texto"
	fmt.Println(str2)

	// '' -> tabela ASCII
	char := 'B'
	fmt.Println(char)

	var erro error = errors.New("Erro interno")
	fmt.Println(erro)


}