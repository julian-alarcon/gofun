package main

import "fmt"

func main() {

	type learningCourse struct {
		name       string
		hours      int
		completion bool
	}

	// Just Declaration without initialization
	// var declaredStruct1 learningCourse
	// declaredStruct2 := new(learningCourse)

	currentOpenCourse := learningCourse{
		name:       "AWS 101",
		hours:      10,
		completion: false,
	}

	fmt.Println(currentOpenCourse)

	fmt.Println("Name of current course:", currentOpenCourse.name)

	currentOpenCourse.completion = true

	fmt.Println(currentOpenCourse)
}
