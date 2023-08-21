package main

import "fmt"


func fatorial(num int) int {
	if num == 0{
		return 1
	}
	return num * fatorial(num - 1)
}

func fibro(n int) int{
	if n == 1 || n == 0 {
		return 1
	} else {
		return fibro(n - 1) + fibro(n - 2) 
	}

 }


func main(){
	fmt.Println("Fatorial de 4 = ", fatorial(4))
	fmt.Println("Fatorial de 4 = ", fatorial(5))
	fmt.Println("Fatorial de 4 = ", fatorial(6))
	fmt.Println("Fatorial de 4 = ", fatorial(7))
	fmt.Println("Fibonacci(7) = ", fibro(7))
	fmt.Println("Fibonacci (4) = ", fibro(4))
	fmt.Println("Fibonacci (4) = ", fibro(1))
	fmt.Println("Fibonacci (4) = ", fibro(2))
	fmt.Println("Fibonacci (4) = ", fibro(5))
}