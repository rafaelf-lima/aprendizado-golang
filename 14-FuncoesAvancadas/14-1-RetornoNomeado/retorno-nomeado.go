package main

import "fmt"


func calculosMatematicos(n1, n2 int) (soma int, subtracao int){
	soma = n1 + n2
	subtracao = n1 - n2
	return
}


func main(){
	r1, r2 := calculosMatematicos(9, 3)
	fmt.Println("Adição: ", r1)
	fmt.Println("Subtração: ", r2)

}