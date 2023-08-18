package main

import (
	"fmt"
	"math"
)


func main(){
	var raio float64
	fmt.Println("Seja bem-vindo!")
	fmt.Print("Informe o raio do círculo: ")
	fmt.Scanln(&raio)
	fmt.Println("Calculando a área...")
	var area float64 = math.Pi * (raio * raio)
	fmt.Printf("A área do círculo com o raio inserido é: %0.2f\n", area)
}