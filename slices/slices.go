package main

import "fmt"

func main() {
	monthToGetCost := make([]string, 3, 4)
	monthToGetCost[0] = "January"
	monthToGetCost[1] = "February"
	monthToGetCost[2] = "March"

	servicesToGetCost := []string{"AWS", "Fastly", "Github"}

	fmt.Printf("Length if the slice is %d and capacity is %d of monthToGetCost\n",
		len(monthToGetCost), cap(monthToGetCost))

	fmt.Printf("Length if the slice is %d and capacity is %d of servicesToGetCost\n",
		len(servicesToGetCost), cap(servicesToGetCost))

	monthToGetCost = append(monthToGetCost, "April")
	for _, month := range monthToGetCost {
		fmt.Println(month)
	}
	fmt.Printf("Length if the slice is %d and capacity is %d of monthToGetCost after adding 1\n",
		len(monthToGetCost), cap(monthToGetCost))

	monthToGetCost = append(monthToGetCost, "June")
	for _, month := range monthToGetCost {
		fmt.Println(month)
	}

	fmt.Printf("Length if the slice is %d and capacity is %d of monthToGetCost after adding a new element beyond the initial defined size\n",
		len(monthToGetCost), cap(monthToGetCost))

	servicesMostUsed := servicesToGetCost[:2]
	fmt.Println(servicesMostUsed)

	newMonths := []string{"May", "June", "July"}
	monthToGetCost = append(monthToGetCost, newMonths...)

	monthToGetCost = append(monthToGetCost, "June")
	for _, month := range monthToGetCost {
		fmt.Println(month)
	}
}
