package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	dict := map[string]int{
		"A X": 1 + 1 + 3,
		"A Y": 1 + 2,
		"A Z": 1 + 3 + 6,
		"B X": 2 + 1 + 6,
		"B Y": 2 + 2 + 3,
		"B Z": 2 + 3,
		"C X": 3 + 1,
		"C Y": 3 + 2 + 6,
		"C Z": 3 + 3 + 3,
	}

	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	games := strings.Split(string(input), "\n")

	score := 0

	for i := range games {
		value, _ := dict[games[i]]
		score += value
	}

	fmt.Printf(fmt.Sprint(score))

}
