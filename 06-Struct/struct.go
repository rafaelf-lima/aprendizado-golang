// Go não tem classe
package main

import "fmt"


type usuario struct {
	nome string
	idade uint8
} 

type endereco struct {
	rua string
	numero uint8
}


func main() {
	//Struct: coleção de campos
	var u usuario
	u.nome = "David"
	u.idade = 21
	fmt.Println(u)

	end := endereco{rua: "Rua Rua", numero: 101}
	fmt.Println(end)



	u2 := usuario{"Davi", 20}
	fmt.Println(u2)


	u3 := usuario{nome: "Davi"}
	fmt.Println(u3)
}