package main

import "fmt"


func recuperarExecucao(){
	// 3
	if r := recover(); r != nil {
		fmt.Println("Execução recuperada com sucesso")
	}
}


func estaAprovado(n1, n2 float64) bool{
	// 2
	defer recuperarExecucao()
	media := (n1 + n2) / 2
	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}
	// 1
	panic("A média é exatamente 6")
}

func main(){
	fmt.Println(estaAprovado(6, 6))
	fmt.Println("Pós execuçaõ")


}