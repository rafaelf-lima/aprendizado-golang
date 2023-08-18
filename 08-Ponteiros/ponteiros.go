package main

import "fmt"

func main(){
	fmt.Println("Ponteiros - Referência de memória")

	var v1 int = 10
	var v2 int = v1

	fmt.Println(v1, v2)

	v1++
	fmt.Println(v1, v2)

	var variavel3 int
	var ponteiro *int

	variavel3 = 100
	ponteiro = &variavel3

	fmt.Println(variavel3, *ponteiro)


	fmt.Println(variavel3, ponteiro)
	
	variavel3 = 150


	fmt.Println(variavel3, ponteiro)

}