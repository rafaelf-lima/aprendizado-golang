package main

import "fmt"

func main() {
	r1 := soma(4,5 )
	fmt.Printf("O resultado é %v", r1)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	r2 := f("Texto de F1")
	fmt.Println(r2)

	rSo, _ := calculosMatematicos(10, 15)
	fmt.Println(rSo)
}



func calculosMatematicos(n1, n2 int8) (int8, int8){
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}


func soma(x, y int) int{
	return x + y

}

