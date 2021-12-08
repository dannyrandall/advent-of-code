package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("../input.txt")
	if err != nil {
		log.Fatalf("unable to open file: %s", err)
	}
	defer f.Close()

	largerThanPrev := 0
	prev := 0
	first := true

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		cur, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatalf("unable to parse number: %s", err)
		}

		if first {
			first = false
			prev = cur
			continue
		}

		if cur > prev {
			largerThanPrev++
		}

		prev = cur
	}

	log.Printf("largerThanPrev: %v", largerThanPrev)
}
