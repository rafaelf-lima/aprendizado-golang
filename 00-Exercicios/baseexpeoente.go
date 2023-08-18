package main

import "fmt"

func init() {
	fmt.Println("Seja bem-vindo!")
}

func main() {
	var base, expoente int
	fmt.Print("Informe a base: ")
	fmt.Scanln(&base)
	fmt.Print("Informe o expoente: ")
	fmt.Scanln(&expoente)
	novoNumero := calculo(base, expoente)
	fmt.Printf("Base (%d) ** Expoente (%d)\n", base, expoente)
	fmt.Printf("O novo número é %d\n", novoNumero)

}

func calculo(b, e int) int{
	var resultado, i int = 1, 0
	for i < e{
		resultado = resultado * 2
		i++
	}
	return resultado
}