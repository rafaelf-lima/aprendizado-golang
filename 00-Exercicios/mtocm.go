package main

import "fmt"



func main(){
	var metros int
	fmt.Println("Seja bem-vindo!")
	fmt.Print("Insira um valor em metros: ")
	fmt.Scanln(&metros)
	fmt.Println("Transformando em centímetros...")

	var centimetros int = metros * 100 
	fmt.Printf("O valor inserido em centímetros é %d\n", centimetros)
}