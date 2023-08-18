package main

import (
	"fmt"
	"time"
)

func main(){
	// Concorrência != Paralelismo
	// Paralelismo: duas ou mais tarefas executadas em mais tarefas (processadores com mais de um núcleo)
	// Concorrência: não necessariamente estão ao mesmo tempo; as tarefas ficariam "revezando" tempo

	// Se colocar go em ambas, não escreve em ambas porque "não há tempo"
	go escrever("Olá, mundo!")
	escrever("Programando em Go!")


}


func escrever(texto string){
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}