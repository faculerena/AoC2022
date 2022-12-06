package main

import (
	"fmt"
	"os"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	sol1, sol2 := marker(string(input), 4), marker(string(input), 14)

	fmt.Println(sol1, sol2)

}

func marker(pre string, pos int) int {
	for i := pos; i < len(pre); i++ {
		set := map[rune]struct{}{}
		for _, r := range pre[i-pos : i] {
			set[r] = struct{}{}
		}
		if len(set) >= pos {
			return i
		}
	}
	return -1
}
