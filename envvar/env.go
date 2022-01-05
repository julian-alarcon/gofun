package main

import (
	"fmt"
	"os"
)

func main() {
	for _, env := range os.Environ() {
		fmt.Println(env)
	}

	name := os.Getenv("USER")
	fmt.Println("My username", name)
}
