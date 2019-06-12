package main

import (
	"fmt"
)

func main() {
	gameScores := make(map[string]int)
	gameScores["WoW"] = 8
	gameScores["Diablo"] = 9
	gameScores["Starcraft"] = 3
	gameScores["Hearthstone"] = 7
	gameScores["COD"] = 6
	gameScores["WarcraftIII"] = 10

	for key, value := range gameScores {
		fmt.Printf("\nKey is %v Value is %v", key, value)
	}
}
