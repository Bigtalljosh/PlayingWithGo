package main

import (
	"fmt"
)

func main() {
	
	type courseMeta struct { 
		Author string
		Level string
		Rating float64
	}

	//var dockerdeepdive courseMeta
	//dockerdeepdive := new(courseMeta)
	dockerdeepdive := courseMeta {
		Author: "Josh",
		Level: "Intermediate",
		Rating: 5,
	}

	fmt.Println(dockerdeepdive)
	fmt.Println("Author is", dockerdeepdive.Author)
}
