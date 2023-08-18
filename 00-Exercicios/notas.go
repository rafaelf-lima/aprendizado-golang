package main

/*

    Faça um programa que peça as três notas da disciplina (AP1, AP2 e AC) e mostre a média. A média é calculada de acordo com a fórmula M = (AP1 + AP2) * 0.4 + AC * 0.2.
    Faça um programa que peça um comprimento em metros e converta para centímetros.
    Faça um programa que leia o raio de um círculo, calcule e mostre sua área. Considere pi como aproximadamente igual a 3.14.
    Monte um conversor de temperatura, que lê uma temperatura em graus Fahrenheit e mostre a temperatura em graus Celsius. A fórmula para conversão é C / 5 = (F - 32) / 9.
*/

import "fmt"


func main(){
	var ap1, ap2, ac, media float64
	fmt.Println("Seja bem-vindo!")
	for { fmt.Print("Insira a primeira nota (AP1): ")
	fmt.Scanln(&ap1)
	if ap1 < 0 || ap1 > 10{
		fmt.Println("Nota inválida. Entre 0 e 10.")
		fmt.Println("Insira novamente a primeira nota (AP1): ")
        fmt.Scanln(&ap1)
	}
	fmt.Print("Insira a segunda nota (AP2): ")
	fmt.Scanln(&ap2)
	if ap2 < 0 || ap2 > 10{
		fmt.Println("Nota inválida. Entre 0 e 10.")
		fmt.Println("Insira novamente a segunda nota (AP2): ")
        fmt.Scanln(&ap2)
	}
	fmt.Print("Insira a nota da avaliação continuada (AC): ")
	fmt.Scanln(&ac)
	if ac < 0 || ac > 10{
		fmt.Println("Nota inválida. Entre 0 e 10.")
		fmt.Println("Insira novamente a nota da avaliação continuada (AC): ")
        fmt.Scanln(&ac)
	}

	media = (ap1 + ap2) * 0.4 + ac * 0.2

	if media > 7 {
		fmt.Printf("Você está aprovado com média %0.2f\n", media)
	} else if media < 7{
		fmt.Printf("Você está reprovado com média %0.2f\n", media)
	}
}
	
	
}