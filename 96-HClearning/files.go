package main

import (
	"fmt"
	"io"
	"os"
)

func main(){
	fmt.Println("Welcome to files in golang")
	content := "This needs to go in a file"

	file, err := os.Create("./mynewfile.txt")
	checkNilErr(err)

	length, err :=io.WriteString(file, content)
	checkNilErr(err)

	fmt.Println("Length: ", length)
	defer file.Close()
	readFile("./mynewfile.txt")
}

func readFile(filname string){
	databyte, err := os.ReadFile(filname)
	checkNilErr(err)

	fmt.Println("Text data inside the file is \n", databyte)
}

func checkNilErr(err error){
	if err != nil{
		panic(err)
	}
}