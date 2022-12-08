package main

import (
	"fmt"
	"os"
	"strings"
)

var input, _ = os.ReadFile("myinput.txt")
var lines = strings.Split(string(input), "\n")
var saved = -1
var sol1 = 0
var mat1 [99][99]int

func main() {
	sol1, _ = part1()
	sol2 := part2()
	fmt.Println(sol1)
	fmt.Println(sol2)

}

func part1() (int, [99][99]int) {
	for k := range lines { //ltr, rtl
		saved = -1
		for i := 0; i < len(lines[k])-1; i++ { //left to right
			if int(lines[k][i]) > saved {
				saved = int(lines[k][i])
				mat1[k][i] = 1
			}
		}
		saved = -1
		for i := len(lines[k]) - 1; i > 0; i-- { //right to left
			if int(lines[k][i]) > saved {
				saved = int(lines[k][i])
				mat1[k][i] = 1
			}
		}
	}
	for k := range lines[1] { //ttb, btt
		saved = -1
		for i := 0; i < len(lines)-1; i++ { //top to bottom
			if int(lines[i][k]) > saved {
				saved = int(lines[i][k])
				mat1[i][k] = 1
			}
		}

		saved = -1

		for i := len(lines) - 2; i > 0; i-- { //bottom to top
			if int(lines[i][k]) > saved {
				saved = int(lines[i][k])
				mat1[i][k] = 1
			}

		}
	}
	for i := 0; i < len(mat1); i++ {
		for j := 0; j < len(mat1); j++ {
			if mat1[i][j] == 1 {
				sol1++
			}
		}
	}

	return sol1, mat1
}

func part2() int {
	sol2 := make([]int, 0)
	_, visible := part1()

	for i := 0; i < len(visible); i++ {
		for j := 0; j < len(visible); j++ {
			if visible[i][j] == 1 {

				//aca tengo que chequear que valor tiene ese [i][j] en lines[i][j] y comparar ese valor con cuantos arboles se ven desde ese para las 4 direcciones

			}
		}
	}

	return higher(sol2)
}

func higher(a []int) int {
	h := 0
	for _, v := range a {
		if v > h {
			h = v
		}
	}
	return h
}
