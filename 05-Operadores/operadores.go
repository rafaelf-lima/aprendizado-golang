package main

import "fmt"

func main(){
	soma := 1 + 2
	subtracao := 10 - 3
	multiplicacao := 5 * 7
	divisao := 2 / 1
	resto := 10 % 2


	fmt.Println(soma, subtracao, multiplicacao, divisao, resto)

	var n1 int16 = 10
	var n2 int16 = 25
	var soma2 int16 = n1 + n2
	fmt.Println(soma2)

	// Atribuição
	var v1 string = "String"
	v2 := "String2"
	fmt.Println(v1, v2)

	// Relacionais
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 > 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)

	// Lógicos
	fmt.Println("-----")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)


	// Unários

	numero := 10
	numero++
	numero += 15
	fmt.Println(numero)

	numero--
	numero-=20
	numero /= 10
	numero %= 3
}
