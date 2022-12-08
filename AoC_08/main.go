package main

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

//aca tengo que chequear que valor tiene ese [i][j] en lines[i][j] y comparar ese valor con cuantos arboles se ven desde ese para las 4 direcciones

func part2() int {
	sol2 := make([][]int, len(lines))

	//explore all 4 directions of each point in "lines" and check what is the distance on each direction to the next higher tree
	for x := 0; x < len(lines); x++ {
		for y := 0; y < len(lines); y++ {
			tree := int(lines[x][y])

			for k := range lines { //ltr, rtl
				count := 0
				for i := x; i < len(lines[k])-1; i++ { //to right
					if int(lines[k][i]) > tree {
						count++
					} else {
						sol2[x][y] *= count

					}
				}

				count = 0
				for i := x; i > 0; i-- { //to left
					if int(lines[k][i]) > tree {
						count++
					} else {
						sol2[x][y] *= count

					}
				}
			}
			for k := range lines[1] { //ttb, btt
				count := 0
				for i := y; i < len(lines)-1; i++ { //to bottom
					if int(lines[k][i]) > tree {
						count++
					} else {
						sol2[x][y] *= count

					}
				}

				count = 0
				for i := y; i > 0; i-- { //to top
					if int(lines[k][i]) > tree {
						count++
					} else {
						sol2[x][y] *= count

					}
				}
			}

		}
	}

	return higher(sol2)

}

func higher(a [][]int) int {

	//find the higher value of the matrix and store it in "higher", and return it
	higher := 0
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a); j++ {
			if a[i][j] > higher {
				higher = a[i][j]
			}
		}
	}
	return higher

}
