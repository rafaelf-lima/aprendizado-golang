package main

import "fmt"


func main(){
	// Make to create slices and maps
	languages := make(map[string]string)

	languages["RB"] = "Ruby"
	languages["PY"] = "Python"
	languages["PHP"] = "PHP"

	for _, value := range languages{
		fmt.Printf("%v", value)
	}

}