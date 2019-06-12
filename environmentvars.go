package main

import (
	"fmt"
	"os"
)

func main() {
	name := os.Getenv("USERNAME")
	course := "My Course"

	fmt.Println("\nHi", name, "you're currently watching", course)
	changeCourse(&course)
	fmt.Println("\nYou are now on the course", course)
}

func changeCourse(course *string) string {
	*course = "My new Course"

	fmt.Println("\nTrying to change your course to", *course)
	return *course
}