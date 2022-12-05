package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

var input, _ = os.ReadFile("input.txt")
var platform = strings.Split(string(input), "\n\n") //devuelve la plataforma y las instrucciones
var boxes = strings.Split(platform[0], "\n")        // devuelve cada altura
var top = boxes[len(boxes)-1]                       //devuelve la ultima altura
var spart1 = map[rune]string{}
var spart2 = map[rune]string{}

type steps struct { //ejemplo = move 7 from 3 to 9
	quantity int
	from     rune
	to       rune
}

func main() {

	part1, part2 := elfDo()
	fmt.Println(part1)
	fmt.Println(part2)
}

func elfDo() (string, string) {

	for _, s := range boxes { //esto arregla la plataforma
		for i, r := range s {
			if unicode.IsLetter(r) { //si es una letra (y no un espacio)
				spart1[[]rune(top)[i]] += string(r)
				spart2[[]rune(top)[i]] += string(r)
			}
		}
	}
	elfMove()
	p1, p2 := elfCountTop()
	return p1, p2
}

func elfMove() {

	for _, s := range strings.Split(strings.TrimSpace(platform[1]), "\n") {

		var instruction steps

		fmt.Sscanf(s, "move %d from %c to %c", &instruction.quantity, &instruction.from, &instruction.to)
		q, t, f := instruction.quantity, instruction.to, instruction.from

		for i := 0; i < q; i++ {
			spart1[t] = spart1[f][:1] + spart1[t]
			spart1[f] = spart1[f][1:]
		}

		spart2[t] = spart2[f][:q] + spart2[t]
		spart2[f] = spart2[f][q:]
	}
}

func elfCountTop() (string, string) {

	part1, part2 := "", ""
	for _, r := range strings.ReplaceAll(top, " ", "") {
		part1 += spart1[r][:1]
		part2 += spart2[r][:1]
	}
	return part1, part2
}
