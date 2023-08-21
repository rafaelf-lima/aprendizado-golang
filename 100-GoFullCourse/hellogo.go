/*

You probably don't need a framework. `net/http` covers most of what you need.

Once you have a concrete problem the standard library can't solve, **then** start looking for the simplest solution you can find for your specific problem.

Estudar sobre o net/http e depois pesquisar web frameworks








*/


package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"golang.org/x/exp/utf8string"
)


var p1 = fmt.Println

func main(){
	/* p1("Hello, Go")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	if err == nil {
		p1("Hello", name)
	} else {
		log.Fatal(err)
	}

	*/

	/* var vName string = "Derek"
	var v1, v2 = 1.2, 3.4
	var v3 = "Hello"
	v4 := */

	/* p1(reflect.TypeOf(25))
	p1(reflect.TypeOf(3.13))
	p1(reflect.TypeOf(true))
	p1(reflect.TypeOf("Hello")) */

	/* cv1 := 1.5
	cv2 := int(cv1)
	p1(cv2) // 1
	cv3 := "5000000"
	cv4, err := strconv.Atoi(cv3)
	p1(cv4, err, reflect.TypeOf(cv4))

	cV7 := "3.14"
	if cv8, err := strconv.ParseFloat(cV7, 64); err == nil 
	{
		p1(cv8)
	}
	cv9 := fmt.Sprintf("%f", 3.14) */

	/* iAge := 8
	if (iAge >= 1) && (iAge <= 18){
		p1("Import birthday")
	} else if (iAge == 21) || (iAge == 50){
		p1("Import birthday")
	} else if (iAge >= 65) {
		p1("Import birthday")
	} else {
		p1("Not an important value")
	}
	p1("!true =", !true)

	sv1 := "A word"
	// array of bytes
	replacer := strings.NewReplacer("A", "Another")
	sv2 := replacer.Replace(sv1)
	p1(sv2)
	p1("Length: ", len(sv2))
	p1("Contains Another: ", strings.Contains(sv2, "Another"))
	p1("o index: ", strings.Index(sv2, "o"))
	p1("Replace: ", strings.Replace(sv2, "o", "0", -1))
	sv3 := "\nSome Words\n"
	sv3 = strings.TrimSpace(sv3)
	p1("Split: ", strings.Split("a-b-c-d", "-"))
	p1(strings.ToLower(sv1))
	p1(strings.ToUpper(sv2))
	p1(strings.HasPrefix("tacocat", "taco"))
	p1(strings.HasSuffix("tacocat", "cat")) */

	/* rStr := "abcdefg"
	p1("Rune count: ", utf8.RuneCountInString(rStr))
	for i, runeVal := range rStr{
		fmt.Printf("%d : %#U : %c\n", i, runeVal, runeVal)
	}

	now := time.Now()
	p1(now.Year(), now.Month(), now.Day())
	p1(now.Hour(), now.Minute(), now.Second())

	for x := 1; x <= 5; x++{
		p1(x)
	}

	fX := 0
	for fX < 5 {
		p1(fX)
		fX++
	}


	aNums := []int{1,2,3}
	for _, num := range aNums {
		p1(num)
	}

	var arr1 [5]int
	arr1[0] = 1
	arr2 := []int{1,2,3,4,5}
	p1(arr2)
	p1(arr1)
	p1(len(arr2))

	for i, v := range arr2 {
		fmt.Printf("%d : %d", i, v )
	}

	arr3 := [2][2]int{
		{1,2},
		{3,4},
	}
	for i:= 0; i < 2; i++{
		for j := 0; j < 2; j++{
			p1(arr3[i][j])
		}
	}
	*/
	/*
	sli1 := make([]string, 6)
	sli1[0] = "Society"
	sli1[1] = "of"
	sli1[2] = "the"
	sli1[3] = "Simulated"
	sli1[4] = "Universe"
	p1(len(sli1))
	for i:=0; i < len(sli1); i++{
		p1(sli1[1])
	}

	for _, x := range sli1 {
		p1(x)
	}


	sArr := [5]int{1,2,3,4,5}
	sl3 := sArr[0:2]
	p1(sl3)

	p1(getSum2(1,2,2,3,3,3,3))
	arrayVar := []int{1,2,3,4}
	p1(getArraySum(arrayVar))



}
// Go permite retrocompatibilidade

func getSum2(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}


func getArraySum(arr []int) int{
	sum := 0
	for _, val := range arr {
		sum += val
	}
	return sum

	*/ 
	// File IO
	/*
	f, err := os.Create("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	iPrimeArr := []int{2,3,5,7,11}
	var sPrimeArr []string
	for _, i := range iPrimeArr{
		sPrimeArr = append(sPrimeArr, strconv.Itoa(i))
	}
	for _, num := range sPrimeArr{
		_, err := f.WriteString(num + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}

	f, err = os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scan1 := bufio.NewScanner(f)
	for scan1.Scan(){
		p1("Prime: ", scan1.Text())
	}
	if err := scan1.Err(); err != nil {
		log.Fatal(err)
	}

	_, err := os.Stat("data.txt")
	if errors.Is(err, os.ErrNotExist){
		p1("File unfound")
	}
	*/
	getSumGen(5,4)
	getSumGen(5.0,4.0)
}

// you cannot define fields in interfaces
// Go interfaces don't allow fields to be declared because conceptually they deal with behavior not data. Fields are data.

type MyConstraint interface {
	int | float64
}

// go mod init nomedoprojeto (criação de um "package.json, no caso o go.mod")


func getSumGen[T MyConstraint](x T, y T) T{
	return x + y
}