package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")

	sol1, sol2 := 0, 0
	for _, tarea := range strings.Fields(string(input)) {
		zonas := strings.Split(tarea, ",")

		elf1 := strings.Split(zonas[0], "-")
		efl2 := strings.Split(zonas[1], "-")

		a, _ := strconv.Atoi(elf1[0])
		b, _ := strconv.Atoi(elf1[1])
		c, _ := strconv.Atoi(efl2[0])
		d, _ := strconv.Atoi(efl2[1])

		if (a >= c && b <= d) || (c >= a && d <= b) {
			sol1++
		}

		if !(b < c || d < a) {
			sol2++
		}
	}
	fmt.Println(sol1)
	fmt.Println(sol2)
}
