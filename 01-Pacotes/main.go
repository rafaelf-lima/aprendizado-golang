package main

import (
	"fmt"
	"modulo/Auxiliar"
	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Olá, mundo")
	// Letra maiúscula: pode ser exportada por outros pacotes
	// Letra minúscula: espécie de private method
	auxiliar.Escrever()
	// O código compilado (go build) não atualiza conforme o código é modificado, deve-se "buildar" novamente
	
	// Usar apenas o último nome pós última barra (/)
	erro := checkmail.ValidateFormat("devbook@gmail.com")
	fmt.Println(erro)

}