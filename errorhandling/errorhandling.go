package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Open("./testfile.txt")

	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println("err value when no error is returned: ", err)
}
