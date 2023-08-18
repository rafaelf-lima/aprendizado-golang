// Métodos estão associados a algo (struct, interface)

package main

import "fmt"

type usuario struct{
	nome string
	idade uint8
}

// Função associada a struct usuário (método)
// u: referenciar outros campos do usuário que chamou (this)
func (u usuario) salvar(){
	fmt.Printf("Salvando os dados do usuário %v no banco de dados", u.nome)
}

func (u usuario) maiorIdade() bool{
	return u.idade >= 18
}

// Uso do ponteiro: alteração muda também fora do método
func (u *usuario) fazerAniversario(){
	u.idade++
}


func main(){
	u1 := usuario{"Usuário 1", 20}
	fmt.Println(u1)
	u1.salvar()
	u1.maiorIdade()
	u2 := usuario{"Davi", 30}
	u2.salvar()
	r1 := u2.maiorIdade()
	fmt.Println(r1)

	u2.fazerAniversario()
	fmt.Println(u2.idade)
}

