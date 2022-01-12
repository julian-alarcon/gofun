package main

import (
	"fmt"
)

func main() {
	// mapUsingMake := make(map[string]int)

	downloadStats := map[string]int{
		"log4j":     14000,
		"simplesum": 3,
		"logback":   500,
	}

	fmt.Println(downloadStats)

	monthDownloadstats := map[string]int{
		"November": 500,
		"December": 800,
		"January":  500,
		"February": 700,
		"March":    800,
	}

	monthDownloadstats["April"] = 600
	fmt.Println(monthDownloadstats)
	for i, j := range monthDownloadstats {
		fmt.Println(i, j)
	}

	delete(monthDownloadstats, "November")
	fmt.Println(monthDownloadstats)
}
