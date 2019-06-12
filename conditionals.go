package main

import (
	"fmt"
)

func main() {
	switch "Test" {
	case "Docker Deep Dive" :
		fmt.Println("bingo")
	case "Test" :
		fmt.Println("nah")
		fallthrough
	default :
		fmt.Println("yeah nah")
	}
}
