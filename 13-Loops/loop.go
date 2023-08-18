package main

import (
	"fmt"
	"time"
)

func main(){
	i := 0

	for i < 10 {
		time.Sleep(1)
		i++
	}

	fmt.Println(i)

	for j := 0; j < 10; j++{
		time.Sleep(time.Second)
		fmt.Println(j)


		nomes := [3]string{"João", "Davi", "Ana"}
		for _, nome := range nomes{
			fmt.Println(nome)
		}

		for indice, letra := range "PALAVRA"{
			fmt.Println(indice, string(letra))
		}



		usuario := map[string]string{ 
			"nome": 	"Leonardo",
			"sobrenome": "Silva",
		}

		for chave, valor := range usuario{
			fmt.Println(chave, valor)
		}

		for {
			fmt.Println("Executar para sempre")
			time.Sleep(time.Second)
		}
	}


}