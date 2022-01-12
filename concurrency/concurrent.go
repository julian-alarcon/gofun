package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitingGrp sync.WaitGroup
	waitingGrp.Add(2) // Setting the number of goroutines to review

	go func() {
		defer waitingGrp.Done()
		time.Sleep(5 * time.Second)
		fmt.Println("Hello now?")
	}()

	func() {
		defer waitingGrp.Done()
		time.Sleep(3 * time.Second)
		fmt.Println("Something in the middle")
	}()

	go func() {
		defer waitingGrp.Done()
		fmt.Println("What about now?")
	}()
	waitingGrp.Wait()
}
