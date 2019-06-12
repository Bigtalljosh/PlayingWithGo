package main

import (
	"fmt"
)

/*
func main() {
	runtime.GOMAXPROCS(2)
	var waitGroup sync.WaitGroup
	waitGroup.Add(2)

	go func() {
		defer waitGroup.Done()

		time.Sleep(5 * time.Second)
		fmt.Println("Hello")
	}()

	go func() {
		defer waitGroup.Done()
		fmt.Println("Pluralsight")
	}()

	waitGroup.Wait()
}*/

func main() {
	myChannel := make(chan int, 5)

	fmt.Println(myChannel)
}