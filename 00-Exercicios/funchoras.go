package main

import "fmt"

func main(){
	var salarioHora float64
	var horasTrabalhadas int
	fmt.Println("Seja bem-vindo!")
	fmt.Print("Insira o número de horas trabalhadas: ")
	fmt.Scanln(&horasTrabalhadas)
	fmt.Print("Insira o salário por hora: ")
	fmt.Scanln(&salarioHora)
	salario := retornaSalario(horasTrabalhadas, salarioHora)
	fmt.Printf("O salário a ser recebido é: R$%0.2f\n", salario)
}


func retornaSalario(horasTrabalhadas int, salario float64) float64{
	var salarioCalculado float64 
	salarioCalculado = float64(horasTrabalhadas) * salario
	return salarioCalculado  
}