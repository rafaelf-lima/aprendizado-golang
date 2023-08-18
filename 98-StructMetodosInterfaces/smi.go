// Go way
package main

import "fmt"

// Interface: contrato
type Veiculo interface{
	buzina()
}


type Pessoa struct{
	nome string
	idade int
	veiculo Veiculo
}

type Carro struct {
	fabricante string
	modelo string
	ano int
}

type Moto struct {
	fabricante string
	modelo string
	ano int
}

func (p Pessoa) andou(){
	fmt.Println(p.nome, "andou")
}

func (c Carro) buzina(){
	fmt.Println(c.fabricante, "peeeeen")
}

func (m Moto) buzina(){
	fmt.Println(m.fabricante, "peeeeen")
}

func main(){

	/* bmw := Carro{
		fabricante: "bmw",
		modelo: "xpto",
		ano: 2010,
	}
	*/

	moto := Moto{
		fabricante: "Moto BMW",
		modelo: "AXFD",
		ano: 2005,
	}

	Wesley := Pessoa{
		nome: "Wesley",
		idade: 23,
		veiculo: moto,
	}
	fmt.Println(Wesley.nome)
	Wesley.andou()

}