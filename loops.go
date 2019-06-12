package main

import (
	"fmt"
)

func main() {
	count := 0
	for {
		fmt.Println("Ping ", count)
		count += 1

		if count > 10 {
			break 
		}
	}

	for i := 0; i < 10; i++ {
		fmt.Println("yolo ", i)
	}

	courses := []string {"DDD", "test", "My course" }

	for _, i := range courses {
		fmt.Println(i)
	}
}
