package main

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
)

// var (
// 	name  = "Lucy"
// 	human = "Julian"
// 	food  = "Croquetas"
// 	age   = "1"
// )

const myConst = 186000

func main() {
	name := "Lucy"
	human := "Julian"
	food := "Croquetas"
	age := "1"
	fmt.Println("Name: ", name)
	fmt.Println("Human: ", human)
	fmt.Println("Food: ", food)
	fmt.Println("Age: ", age)
	fmt.Println("Type of name variable ", reflect.TypeOf((name)))
	fmt.Println("Type of age variable ", reflect.TypeOf((age)))
	var ptr *string = &age
	intAge, err := strconv.Atoi(*ptr)
	if err == nil {
		fmt.Println("Age:", age, "Converted age ", intAge, reflect.TypeOf(intAge), "\nError: ", err)
	} else {
		// myConst = 0
		fmt.Println("Error: ", err, myConst)
		os.Exit(1)
	}
}
