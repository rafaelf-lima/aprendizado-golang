package main

import "fmt"


func main(){
	var qntdEntradas int
	var media float64
	var soma, nums int = 0, 0
	fmt.Println("Olá! Quantos números?")
	fmt.Scan(&qntdEntradas)

		for i := 0; i < qntdEntradas; i++ {
		fmt.Println("Insira um número: ")
		fmt.Scan(&nums)
		soma += nums
	}

	media = float64(soma) / float64(nums)
	fmt.Printf("Média: %0.2f\n",  media)
}