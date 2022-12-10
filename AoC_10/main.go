package main

import (
	"fmt"
	"os"
	"strings"
)

const width = 40

var clock, part1 int
var memvalue int = 1
var part2 string
var input, _ = os.ReadFile("input.txt")
var command, value = "", 0

func main() {

	for _, s := range strings.Split(string(input), "\n") {

		fmt.Sscanf(s, "%s %d", &command, &value)

		cycle()

		if command == "addx" {
			addx(value)
		}
	}

	solution()

}

func cycle() {
	part2 += map[bool]string{true: "##", false: "  "}[clock%width >= memvalue-1 && clock%width <= memvalue+1]
	part2 += map[bool]string{true: "\n"}[clock%width == width-1]
	if clock++; (clock+width/2)%width == 0 {
		part1 += clock * memvalue
	}
}

func addx(value int) {
	cycle()
	memvalue += value
}

func solution() {
	fmt.Println(part1)
	fmt.Println(part2)
}
