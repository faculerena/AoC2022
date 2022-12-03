package main

import (
	"log"
	"os"
	"time"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	time1 := time.Now()

	// part 1 content

	elapsed := time.Since(time1)
	log.Printf("Parte 1 - %s", elapsed)

	time2 := time.Now()

	// part 2 content

	elapsed = time.Since(time2)
	log.Printf("Parte 2 - %s", elapsed)
}
