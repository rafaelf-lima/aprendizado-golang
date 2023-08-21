// Tentativa de fazer a AP1 de Arq. de Computadores usando a linguagem Go.

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
	var resultado string = ""
	for i := 0; i < len(operando); i++{
		if operando[i] == '0' {
			resultado += "1"
		} else {
			resultado += "0"
		}
	
	}
	x = resultado
	// fmt.Println(x)
}

func adicao(input string) {
	primeiraParte := input[3:7]
	segundaParte := input[7:11]
	resultado := binarioDecimal(primeiraParte) + binarioDecimal(segundaParte)
	x = decimalBinario(resultado)
	
}

func adicaoImediata(input string) {
	operandoInput := input[3:11]
	fmt.Println(x)
	resultado := binarioDecimal(operandoInput) + binarioDecimal(x)
	fmt.Println(resultado)
}

func emprestar(minuendo *[]byte, index int) {
	for index >= 0 && (*minuendo)[index] == '0' {
		(*minuendo)[index] = '1'
		index--
	}
	if index >= 0 {
		(*minuendo)[index] = '0'
	}
}

func subtracao(input string) {
	primeiraParte := []byte(input[3:7])
	segundaParte := []byte(input[7:11]) 

	if string(segundaParte) > string(primeiraParte) {
		primeiraParte, segundaParte = segundaParte, primeiraParte
		emprestar(&primeiraParte, len(primeiraParte)-1)
	}

	resultado := binarioDecimal(string(primeiraParte)) - binarioDecimal(string(segundaParte))
	x = decimalBinario(resultado)
}

func subtracaoImediata(input string) {
	fmt.Println("Executing subtracaoImediata")
}

func leitura(input string) {
	operandoInput := input[3:11]
	if operandoInput == "00000000" {
		exibeResultado(x)
	} else {
		exibeResultado(y)
	}
}

func armazena(input string) {
	operandoInput := input[3:11]
	x = operandoInput
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
				fmt.Println("Operação: Complemento de 1")
				complemento1(input)
			case "001":
				fmt.Println("Operação: Adição")
				adicao(input)
			case "010":
				fmt.Println("Operação: Adição imediata")
				adicaoImediata(input)
			case "011":
				fmt.Println("Operação: Subtração")
				subtracao(input)
			case "100":
				fmt.Println("Operação: Subtração imediata")
				subtracaoImediata(input)
			case "101":
				fmt.Println("Operação: Leitura")
				leitura(input)
			case "110":
				fmt.Println("Operação: Armazena")
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
			fmt.Println("Podemos começar? Digite 0 para sair e 1 para iniciar")
			fmt.Scanln(&inicio)
			switch inicio {
			case 0:
				fmt.Println("Até logo!")
				os.Exit(0)
			case 1:
				for {
				var input string
				fmt.Println("Vamos começar! Digite 11 dígitos binários!")
				fmt.Scanln(&input)
				if len(input) == 11 {
					verificaBinarios(input)
				} else {
					fmt.Println("Parece que não funcionou. Tente novamente!")
				}
			}
			default:
				fmt.Println("Opção inválida. Por favor, digite 0 ou 1.")
			}
		}
