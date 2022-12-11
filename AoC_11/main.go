package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

const TURNS = 20

func main() {

	rawinput, _ := ioutil.ReadFile("input.txt")
	input := strings.Split(string(rawinput), "\r\n")

	monkeys := make(map[int][]uint64)

	var monke int
	var sign string
	var value uint64
	var prev bool
	var test uint64
	var catchYes int
	var catchNo int

	var turnList [8]int

	// Part Two
	// lcm := uint64(19 * 3 * 13 * 7 * 5 * 11 * 17 * 2)

	for i := 0; i < TURNS; i++ {
		monke = 0

		for i, line := range input {

			if i == 0 {
				if i == 1+monke*7 {
					regex := regexp.MustCompile(`[0-9]+`)
					items := regex.FindAllString(line, -1)
					for _, item := range items {
						item, _ := strconv.ParseUint(item, 10, 64)
						monkeys[monke] = append(monkeys[monke], item)

					}
					continue
				}
			}
			if i == 2+monke*7 {
				l := strings.Split(line, " ")
				sign = l[6]
				if l[7] == "old" {
					prev = true
				} else {
					prev = false
				}

				if !prev {
					value, _ = strconv.ParseUint(l[8], 10, 64)
				}
			}
			if i == 3+monke*7 {
				l := strings.Split(line, " ")
				test, _ = strconv.ParseUint(l[5], 10, 64)
				continue
			}
			if i == 4+monke*7 {
				l := strings.Split(string(line), " ")
				catchYes, _ = strconv.Atoi(l[9])

				continue
			}

			if i == 5+monke*7 {
				l := strings.Split(string(line), " ")
				catchYes, _ = strconv.Atoi(l[9])
				continue
			}

			if len(line) == 0 {
				for _, worryLevel := range monkeys[monke] {
					turnList[monke]++

					if sign == "+" {
						worryLevel += value
					} else if sign == "*" && !prev {
						worryLevel *= value
					} else if sign == "*" && prev {
						worryLevel *= worryLevel
					}

					worryLevel = worryLevel / 3
					// worryLevel = worryLevel % lcm

					var receiver int
					if worryLevel%test == 0 {
						receiver = catchYes
					} else {
						receiver = catchNo
					}

					monkeys[monke] = monkeys[monke][1:]
					monkeys[receiver] = append(monkeys[receiver], worryLevel)

				}

				monke++
				continue
			}

		}

	}

	countSlice := turnList[:]
	sort.Ints(countSlice)
	fmt.Println("Part One Answer:", countSlice[len(countSlice)-1]*countSlice[len(countSlice)-2])
	// fmt.Println("Part Two Answer:", countSlice[len(countSlice)-1]*countSlice[len(countSlice)-2])

}
