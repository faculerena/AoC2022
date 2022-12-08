package test

import (
	"fmt"
	"os"
	"strings"
)

var input, _ = os.ReadFile("input.txt")
var lines = strings.Split(string(input), "\n")
var saved = -1
var sol1 = 0
var mat1 [99][99]int

func test() {
	sol1, _ = parte1()
	//	sol2 := part2()
	fmt.Println(sol1)
	//	fmt.Println(sol2)

}

func parte1() (int, [99][99]int) {
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
