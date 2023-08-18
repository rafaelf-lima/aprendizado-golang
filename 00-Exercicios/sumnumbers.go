package main

import "fmt"


func main(){
	var num int
	fmt.Println("Entre com um número: ")
	fmt.Scan(&num)

	var soma = (num * (num + 1)) / 2
	fmt.Printf("Soma: %d\n", soma)

}