package main

import "fmt"


func main(){

	fmt.Println("Welcome to a class on pointers")

	//var ptr *int
	//fmt.Println("Valuje of pointers: ", ptr)

	myNumber := 23
	var ptr *int = &myNumber

	fmt.Println(ptr)
	fmt.Println(*ptr)


	*ptr = *ptr * 2
	fmt.Println("new", myNumber)

}