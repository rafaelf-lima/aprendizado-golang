package main


import (
	"fmt"
	"errors"
)


func main(){
	err := someUncertainOperation()
	if err != nil {
		fmt.Println("Failed")
	}

}
// Check if err != only way (?)

func someUncertainOperation() error {
	return errors.New("Uncertain error")
}