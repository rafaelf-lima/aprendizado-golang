// Dicionário do Python?
package main

import "fmt"

func main(){
	fmt.Println("Maps")

	//[tipoEsq]tipoDir
	usuario := map[string]string{
		"Nome": "Rafael",
		"Idade": "22",
	}

	usuario2 := map[int]string{
		1: "Rafael",
		2: "22",
	}

	usuario3 := map[string]map[string]string{
		"nome":{ 
			"primeiro": "João",
			"ultimo": "Neves",
		},
		"curso": {
			"nome": "Engenharia",
			"campus": "Campus 1",
		},

	}

	fmt.Println(usuario)
	fmt.Println(usuario["Nome"])

	fmt.Println(usuario2)
	fmt.Println(usuario2[1])

	fmt.Println(usuario3)
	fmt.Println(usuario3["nome"])

	delete(usuario2, 1)

	usuario4 := map[string]string{
		"Nome": "Gêmeos",
	}

	fmt.Println(usuario4)

}