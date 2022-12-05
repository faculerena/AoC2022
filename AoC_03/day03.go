package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	time1 := time.Now()

	items := strings.Split(string(input), "\n")
	sum1 := 0
	for _, i := range items {
		middle := len(i) / 2
		fst, snd, m := i[:middle], i[middle:], make(map[rune]bool)
		for _, j := range fst {
			m[j] = false
		}
		for _, j := range snd {
			if end, ok := m[j]; ok && !end {
				m[j] = true
				sum1 += letterValue(j)
			}
		}
	}

	fmt.Println(sum1)
	elapsed := time.Since(time1)
	log.Printf("Parte 1 - %s", elapsed)

	time2 := time.Now()
	sum2 := 0
	for i := 0; i < (len(items) - 2); i += 3 {
		m := make([]int, 53)
		fst, snd, trd := items[i], items[i+1], items[i+2]
		for _, a := range fst {
			lval := letterValue(a)
			if m[lval] == 0 {
				m[lval] = 1
			}
		}
		for _, a := range snd {
			lval := letterValue(a)
			if m[lval] == 1 {
				m[lval] = 2
			}
		}
		for _, a := range trd {
			lval := letterValue(a)
			if m[lval] == 2 {
				sum2 += lval
				break
			}
		}
	}
	fmt.Println(sum2)

	elapsed = time.Since(time2)
	log.Printf("Parte 2 - %s", elapsed)

}

func letterValue(c rune) int {
	if c <= 'Z' {
		return int(c - 'A' + 27)
	}
	return int(c - 'a' + 1)
}
