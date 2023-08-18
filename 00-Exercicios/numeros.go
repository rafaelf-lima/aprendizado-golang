package main

import (
	"fmt"
	"math"
)

func init() {
	fmt.Println("Seja bem-vindo!")
}

func main() {
	var num1, num2, num3 int
	fmt.Print("Informe o número 1: ")
	fmt.Scanln(&num1)
	fmt.Print("Informe o número 2: ")
	fmt.Scanln(&num2)
	fmt.Print("Informe o número 3: ")
	fmt.Scanln(&num3)

	novoNum1,novoNum2, novoNum3  := manipulaNumeros(num1, num2, num3)
	fmt.Printf("O primeiro número agora é %d, o segundo número é %d e o terceiro número agora está com %d\n", novoNum1, novoNum2, novoNum3)

}


func manipulaNumeros(n1, n2, n3 int) (int, int, int){
	nN1 := (n1 * 2) * (n2 / 2)
	nN2 := (n1 * 3) + n3
	nN3 := int(math.Pow(float64(n3), 3))
	return nN1, nN2, nN3
}