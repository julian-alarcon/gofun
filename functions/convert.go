package main

import (
	"fmt"
	"strings"
)

func main() {
	name := "julian alarcon"
	product := "lata de pollo"
	fmt.Println("Before changes: ", name, product)
	converter(&name, &product)
	fmt.Println("After changes: ", name, product)
}

func converter(name1 *string, product2 *string) {
	*name1 = strings.ToUpper(*name1)
	*product2 = strings.Title(*product2)
}
