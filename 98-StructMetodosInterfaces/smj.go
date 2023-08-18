package main

import "fmt"

type Cliente struct{
	nome string
	email string
	CPF int
}


func (c Cliente) Exibe(){
	fmt.Println("Exibindo cliente pelo métood: ", c.nome)
}

type ClienteInternacional struct {
	Cliente
	pais string
}




func main(){

	cliente := Cliente{"Wesley", "w@w.com", 1212121212}

	fmt.Println(cliente.nome)


	cliente2 := Cliente{"Mari", "m@m.com", 912192192}

	fmt.Printf("O e-mail é %s e o CPF é %d, \n", cliente2.email, cliente2.CPF)

	cliente3 := ClienteInternacional{
		Cliente: Cliente{
			"Davi", "d@d.com", 1212121212,
		},
	}

	fmt.Printf("Nome %s, CPF: %d, país: %s\n", cliente3.nome, cliente3.CPF, cliente3.pais)

	cliente.Exibe()
	cliente2.Exibe()

}