package main

import (
	"fmt"
)

func main() {
	//myCourses := make([]string, 5, 10)
	myCourses := []int{ 1, 2, 3, 4, 5 }

	fmt.Printf("Length is: %d.\nCapacity is: %d.", len(myCourses), cap(myCourses))
}
