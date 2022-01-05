package main

import "fmt"

func main() {
	fmt.Println("Not again...")
breakPoint1:
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		if i == 2 {
			break breakPoint1
		}
	}

	for i := 30; i > 0; i-- {
		if i%3 == 0 {
			fmt.Println("Countdown number is divisible by 3: ", i)
			// break breakPoint1 # This will not work
		}
	}
}
