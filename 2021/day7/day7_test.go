package day7

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestPartOne(t *testing.T) {
	crabs := parseInput(t)

	max := 0
	for _, crab := range crabs {
		if crab > max {
			max = crab
		}
	}

	minCost := 999999999

	for pos := 0; pos <= max; pos++ {
		cost := costToAlign(crabs, pos)
		if cost < minCost {
			minCost = cost
		}
	}

	t.Logf("crabs: %v", minCost)
}

func parseInput(t *testing.T) []int {
	f, err := os.Open("input.txt")
	if err != nil {
		t.Fatalf("unable to open file: %s", err)
	}
	defer f.Close()

	var crabs []int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), ",")
		for _, s := range split {
			n, _ := strconv.Atoi(s)
			crabs = append(crabs, n)
		}

		// only one line
		break
	}

	return crabs
}

func costToAlign(crabs []int, pos int) int {
	cost := 0
	for _, crab := range crabs {
		dist := crab - pos
		if dist < 0 {
			dist *= -1
		}

		cost += dist
	}

	return cost
}

func TestPartTwo(t *testing.T) {
	crabs := parseInput(t)

	max := 0
	for _, crab := range crabs {
		if crab > max {
			max = crab
		}
	}

	minCost := 999999999

	for pos := 0; pos <= max; pos++ {
		cost := costToAlignPartTwo(crabs, pos)
		if cost < minCost {
			minCost = cost
		}
	}

	t.Logf("crabs: %v", minCost)
}

func costToAlignPartTwo(crabs []int, pos int) int {
	cost := 0
	for _, crab := range crabs {
		cost += costToMove(crab, pos)
	}

	return cost
}

func costToMove(curPos, endPos int) int {
	dist := curPos - endPos
	if dist < 0 {
		dist *= -1
	}

	cost := 0
	for i := 1; i <= dist; i++ {
		cost += i
	}

	return cost
}
