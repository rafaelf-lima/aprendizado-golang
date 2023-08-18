package main

import "fmt"


func main() {
	var variavel1 string = "Variável 1"
	variavel2 := "Variável 2"
	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var (
		variavel3 string = "Variável 4"
		variavel4 string = "Variável 5"
	)

	fmt.Println(variavel3)
	fmt.Println(variavel4)

	variavel5, variavel6 := "Variável 5", "Variável 6"

	fmt.Println(variavel5)
	fmt.Println(variavel6)

	const constante1 string = "Constante 1"
	fmt.Println(constante1)

	variavel4, variavel5 = variavel5, variavel4
	




}