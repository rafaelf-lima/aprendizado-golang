package main

import "fmt"

func f1(){
	fmt.Println("Executando a função 1")
}

func f2(){
	fmt.Println("Executando a função 2")
}

func estaAprovado(n1, n2 int) bool{
	defer fmt.Println("Média calculada. Resultado retornado.")
	fmt.Println("Entrando na função")

	media := (n1 + n2) / 2
	if media >= 6 {
		return true
	}

	return false
}

func main(){
	// Defer: adiar
	defer f1()
	f2()

	fmt.Println(estaAprovado(7, 8))


}