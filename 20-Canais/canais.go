package main

import (
	"fmt"
	"time"
)

func main(){
	// Canal de comunicação para envio e recebimento de dados para sincronizar as GoRoutines
	canal := make(chan string)
	go escrever("Olá, mundo", canal)

	// Esperar receber o valor e salvar na variável mensagem
	mensagem := <- canal
	fmt.Println(mensagem)
}


func escrever(texto string, canal chan string){
	for i := 0; i < 5; i++{
		// Mandando um valor para o canala
		canal <- texto
		time.Sleep(time.Second)
	}
}