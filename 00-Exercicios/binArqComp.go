package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var y, x string

func init(){
	fmt.Println("Seja bem-vindo!")
}

func exibeResultado(resultadoFinal string){
	fmt.Printf("O resultado é: %s\n", resultadoFinal)
	fmt.Println(" ")
}

func binarioDecimal(input string) int {
	res, err := strconv.ParseInt(input, 2, 64)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return int(res)
}

func decimalBinario(input int) string {
	res := strconv.FormatInt(int64(input), 2)
	return res
}

func complemento1(input string){
	operando := input[3:11]
	var resultado string = " "
	for i := 0; i < len(operando); i++{
		if operando[i] == '0' {
			resultado += "1"
		} else {
			resultado += "0"
		}
	
	}
	exibeResultado(resultado)
}

func adicao(input string) {
	var resultado int
	primeiraParte := input[3:7]
	segundaParte := input[7:11]
	resultado = binarioDecimal(primeiraParte) + binarioDecimal(segundaParte)
	exibeResultado(decimalBinario(resultado))
	
}

func adicaoImediata(input string) {
	
}

func subtracao(input string) {
	fmt.Println("Executing subtracao")
}

func subtracaoImediata(input string) {
	fmt.Println("Executing subtracaoImediata")
}

func leitura() {
	operandoX := x[3:11]
	if strings.ContainsAny(operandoX, "0"){
		fmt.Println(x)
	} else {
		fmt.Println(y)
	}
}

func armazena(input string) {
	x = input

}

func transfere() {
	y = x
}


func verificaBinarios(input string) {
	reinicia := false
	fmt.Println("Processando a entrada de bits...")
	if strings.ContainsAny(input, "01") {
			defineOp := input[:3]
			switch defineOp {
			case "000":
				complemento1(input)
			case "001":
				adicao(input)
			case "010":
				adicaoImediata(input)
			case "011":
				subtracao(input)
			case "100":
				subtracaoImediata(input)
			case "101":
				leitura()
			case "110":
				armazena(input)
			case "111":
				transfere()
			default:
				fmt.Println("Você inserou dígitos diferentes de 0 e 1!")
				fmt.Println(" ")
				reinicia = true
			}
		} 
		if reinicia {
			main()
		}
	}


func main(){
		var inicio int
		var input string
		for { 
			fmt.Println("Podemos começar? Digite 0 para sair e 1 para iniciar")
			fmt.Scanln(&inicio)
		if inicio == 1 {
			fmt.Println("Vamos começar! Digite 11 dígitos binários!")
			fmt.Scanln(&input)
			if len(input) == 11 {
				verificaBinarios(input)
			} else {
				fmt.Println("Parece que não funcionou. Tente novamente!")
			}
		} else if inicio == 0 {
			os.Exit(1)
		} else{ 
			fmt.Println("Aparentemente você recusou o inicío. Aperte qualquer botão para começar!")
			fmt.Scanln(&inicio)
		}
	}

	}