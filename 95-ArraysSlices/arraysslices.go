package main

import "fmt"

func main(){
	var arr1[3]int
	fmt.Println(arr1)
	fmt.Println(arr1[0])

	var arr3 = [3]int{2,2,2}
	fmt.Println(arr3)

	var arr2 [4]int = [4]int {4,4,2,1}
	fmt.Println(arr2)

	arr4 := [4]int{2,2,2,2}
	fmt.Println(arr4)
	fmt.Println(arr2 == arr4)

	arr5 := [...]int{5: -1}
	fmt.Println(arr5)


	fmt.Printf("len %d cap %d\n", len(arr1), cap(arr1))


	months := [...]string{
		"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Ago", "Sep", "Oct", "Nov", "Dez",
	}

	summer := months[5:8]
	fmt.Printf("Summer %v\n", summer)
	q2 := months[3:6]
	fmt.Println(q2)
	all := months[:]
	fmt.Println(all)

	slice := make([]int, 2)
	slice[0] = 1
	slice[1] = 2
	printSlice(slice)

	for i := 0; i < len(slice); i++{
		fmt.Printf("slice[%d], %d\n", i, slice[i])
	}

	for _, value := range slice{
		fmt.Print(value)
	}


}


func printSlice(slice []int){
	fmt.Printf("Slice %v len ", slice)
}