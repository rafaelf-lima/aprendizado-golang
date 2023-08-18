package main

import "fmt"

// Go supports anonymous functions, which can form closures. 

// An anonymous function is a function without a name.

func main(){


	r1 := func(texto string)string{
		// Sprintf: formats according to a format specifier and returns the resulting string
		return fmt.Sprintf("Recebido -> %s", texto)
	}("Passando parâmetro")

	fmt.Println(r1)
}