package main

import "fmt"

func soma(numeros ... int) int{
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total

}

func escrever(texto string, numeros ... int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}

}


func main(){
	r1 := soma(1,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2)
	fmt.Println(r1)
	escrever("Olá", 10, 9, 8, 7)
	


}