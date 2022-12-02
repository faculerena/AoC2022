package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	dict := map[string]int{
		"A X": 1 + 3, // piedra piedra
		"A Y": 2 + 6, // piedra papel
		"A Z": 3,     // piedra tijera
		"B X": 1,     // papel piedra
		"B Y": 2 + 3, // papel papel
		"B Z": 3 + 6, // papel tijera
		"C X": 1 + 6, // tijera piedra
		"C Y": 2,     // tijera papel
		"C Z": 3 + 3, // tijera tijera
	}
	dict2 := map[string]int{
		"A X": 3,
		"A Y": 1 + 3,
		"A Z": 2 + 6,
		"B X": 1,
		"B Y": 2 + 3,
		"B Z": 3 + 6,
		"C X": 2,
		"C Y": 3 + 3,
		"C Z": 1 + 6,
	}

	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// part 1
	games, score1, score2 := strings.Split(string(input), "\n"), 0, 0
	for i := range games {
		value_part1, _ := dict[games[i]]
		value_part2, _ := dict2[games[i]]
		score1 += value_part1
		score2 += value_part2
	}
	fmt.Println(fmt.Sprint(score1))
	fmt.Println(fmt.Sprint(score2))

}
