package main

import "fmt"

func main() {
	longestWord := checkLength("isjdg", "absurd", "absurdo", "complicado", "kompliziert")
	fmt.Println("Longest word is ", longestWord)
}

func checkLength(words ...string) string {
	longest := words[0]
	for _, i := range words {
		if len(i) > len(longest) {
			longest = i
		}
	}
	return longest
}
