package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	dict := map[string]int{"A X": 4, "A Y": 8, "A Z": 3, "B X": 1, "B Y": 5, "B Z": 9, "C X": 7, "C Y": 2, "C Z": 6}
	dict2 := map[string]int{"A X": 3, "A Y": 4, "A Z": 8, "B X": 1, "B Y": 5, "B Z": 9, "C X": 2, "C Y": 6, "C Z": 7}
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	games, score1, score2 := strings.Split(string(input), "\n"), 0, 0
	for i := range games {
		value_part1, value_part2 := dict[games[i]], dict2[games[i]]
		score1 += value_part1
		score2 += value_part2
	}
	fmt.Println(fmt.Sprint(score1) + ", " + fmt.Sprint(score2))
}
