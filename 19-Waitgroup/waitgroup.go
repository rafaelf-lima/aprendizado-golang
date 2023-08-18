package main

import (
	"fmt"
	"time"
	"sync"
)

func main(){
	// 
	var waitGroup sync.WaitGroup

	// Quantidade de GoRoutines, "lista de espera"
	waitGroup.Add(2)

	go func(){
		escrever("Olá, mundo")
		// Tira do contador
		waitGroup.Done() // -1 
	}()

	go func(){
		escrever("Programando em Go")
		waitGroup.Done() // -1
	}()

	// Espera até que seja 0
	waitGroup.Wait()


}


func escrever(texto string){
	for i := 0; i < 5; i++{
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}