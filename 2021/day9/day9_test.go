package day9

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"testing"
)

func parseInput(t *testing.T) [][]int {
	f, err := os.Open("input.txt")
	if err != nil {
		t.Fatalf("unable to open file: %s", err)
	}
	defer f.Close()

	var grid [][]int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		var row []int

		for _, r := range scanner.Text() {
			n, _ := strconv.Atoi(string(r))
			row = append(row, n)
		}

		grid = append(grid, row)
	}

	return grid
}

func isLowPoint(grid [][]int, row, col int) bool {
	val := grid[row][col]

	switch {
	// up
	case row > 0 && val >= grid[row-1][col]:
		return false
	// down
	case row < len(grid)-1 && val >= grid[row+1][col]:
		return false
	// left
	case col > 0 && val >= grid[row][col-1]:
		return false
	// right
	case col < len(grid[0])-1 && val >= grid[row][col+1]:
		return false
	}

	return true
}

func TestPartOne(t *testing.T) {
	grid := parseInput(t)

	risk := 0

	// grid[row][col] is correct
	for row := range grid {
		for col := range grid[row] {
			if isLowPoint(grid, row, col) {
				risk += 1 + grid[row][col]
			}
		}
	}

	t.Logf("Risk: %v", risk)
}

type Point struct {
	Row int
	Col int
}

// find all low points, traverse out from each
func TestPartTwo(t *testing.T) {
	grid := parseInput(t)

	var lowPoints []Point

	for row := range grid {
		for col := range grid[row] {
			if isLowPoint(grid, row, col) {
				lowPoints = append(lowPoints, Point{
					Row: row,
					Col: col,
				})
			}
		}
	}

	var sizes []int
	for _, p := range lowPoints {
		basin := make(map[Point]bool)
		exploreBasin(grid, p, basin)
		sizes = append(sizes, len(basin))
	}

	t.Logf("Low Points: %v", lowPoints)
	t.Logf("Sizes : %v", sizes)

	sort.Ints(sizes)

	total := sizes[len(sizes)-3]
	total *= sizes[len(sizes)-2]
	total *= sizes[len(sizes)-1]

	t.Logf("Total: %v", total)
}

func exploreBasin(grid [][]int, p Point, seen map[Point]bool) {
	if seen[p] {
		return
	}

	seen[p] = true

	// up
	if p.Row > 0 && grid[p.Row-1][p.Col] < 9 {
		exploreBasin(grid, Point{p.Row - 1, p.Col}, seen)
	}

	// down
	if p.Row < len(grid)-1 && grid[p.Row+1][p.Col] < 9 {
		exploreBasin(grid, Point{p.Row + 1, p.Col}, seen)
	}

	// left
	if p.Col > 0 && grid[p.Row][p.Col-1] < 9 {
		exploreBasin(grid, Point{p.Row, p.Col - 1}, seen)
	}

	// right
	if p.Col < len(grid[0])-1 && grid[p.Row][p.Col+1] < 9 {
		exploreBasin(grid, Point{p.Row, p.Col + 1}, seen)
	}
}
