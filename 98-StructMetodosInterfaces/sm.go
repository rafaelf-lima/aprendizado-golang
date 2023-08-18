// Go way
package main

import "fmt"

type Pessoa struct{
	nome string
	idade int
	carro Carro

}

type Carro struct {
	fabricante string
	modelo string
	ano int
}

func (p Pessoa) andou(){
	fmt.Println(p.nome, "andou")
}

func (p *Pessoa) mudou(){
	p.nome = "asdasddsa"
	fmt.Println(p.nome, "andou")
}


func main(){

	bmw := Carro{
		fabricante: "bmw",
		modelo: "xpto",
		ano: 2010,
	}

	Wesley := Pessoa{
		nome: "Wesley",
		idade: 23,
		carro: bmw,
	}
	fmt.Println(Wesley.nome)
	Wesley.andou()
	Wesley.mudou()
	fmt.Println(Wesley.nome)



}