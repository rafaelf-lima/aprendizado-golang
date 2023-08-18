package main

import "fmt"

type Carro struct {
	nome string
}

func (c *Carro) andou(){
	c.nome = "BMW"
	fmt.Println(c.nome, "andou")
}



func main(){
	carro := Carro{
		nome: "Ka",
	}

	carro.andou()
	fmt.Println(carro.nome)

/* a := 10

	// Endereço de memória
	fmt.Println(&a)

	// variável do tipo *int - ponteiro sabe onde está localizado o endereço da memória de a
	var ponteiro *int = &a
	fmt.Println(ponteiro)

	// Dereference (caminho inverso para obter o valor do endereço da memória)
	fmt.Println(*ponteiro)


	// Atribui 50 como valor ao endereço da memória que estava 10
	// 10 e 50 apontam pro mesmo endereço
	*ponteiro = 50 
	fmt.Println(*ponteiro)

	// O resultado fica 50 porque altewra o endereçamento da memória
	fmt.Println(a)

	b := &a 
	fmt.Println(b)
	fmt.Println(*b)
	*b = 60
	fmt.Println(a)

	// Variável, endereço de memória e ponteiro

	
	variavel := 10
	abc(&variavel)
	fmt.Println(variavel)

	*/
}


/*

func abc( a *int) {
	*a = 200
}

*/