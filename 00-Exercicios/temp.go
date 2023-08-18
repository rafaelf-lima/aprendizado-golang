package main

import (
	"fmt"
	"time"
)


func main(){
	var fahreint float64
	var celsius float64 = 0
	fmt.Println("Seja bem-vindo!")
	fmt.Print("Informe uma temperatura em Fahrenheit: ")
	fmt.Scanln(&fahreint)
	fmt.Println("Transformando em Celsius...")
	time.Sleep(time.Second)
	celsius = 5 * (fahreint - 32)/9 
	fmt.Printf("A temperatura informada em celsius é %0.2f\n", celsius)

}