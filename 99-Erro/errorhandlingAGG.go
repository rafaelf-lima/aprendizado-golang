package main

import (
	"fmt"
	"os"
	"strconv"
)

// own custom error
type CustomError struct {
	Message string
	Code int
}

func (c CustomError) Error() string{
	return c.Message + " " + strconv.Itoa(c.Code)
}

func Divide(x, y float64) (float64, error){
	if y == 0{
		// return float64(0), errors.New("Can not divide by 0\n")
		return 0.0, CustomError{"Can not divide by 0", -1}
	} else {
		return x / y, nil
	}
}


func main(){
	file, fileErr := os.ReadFile("file")
	if fileErr != nil {
		fmt.Println("Error reading file: ", fileErr)
	} else {
	fmt.Println(file)
	}

	value, divErr := Divide(7,0)
	if divErr != nil {
		fmt.Println(divErr)
	} else {
		fmt.Println(value)
	}

	/* value, _ := Divide(7,0)
	if divErr != nil {
		fmt.Println(divErr)
	} else {
		fmt.Println(value)
	} */
}