package day4

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"testing"
)

type BingoSquare struct {
	Number int
	Marked bool
}

// board[row][col]
type Board [][]BingoSquare

func (b Board) HasBingo() bool {
	// check row
	for row := range b {
		bingo := true
		for col := range b[row] {
			if !b[row][col].Marked {
				bingo = false
				break
			}
		}

		if bingo {
			return true
		}
	}

	// check columns
	for col := 0; col < len(b[0]); col++ {
		bingo := true
		for row := 0; row < len(b); row++ {
			if !b[row][col].Marked {
				bingo = false
				break
			}
		}

		if bingo {
			return true
		}
	}

	return false
}

func (b Board) Mark(num int) {
	for row := range b {
		for col := range b[row] {
			if b[row][col].Number == num {
				b[row][col].Marked = true
			}
		}
	}
}

func (b Board) SumUnmarked() int {
	sum := 0
	for row := range b {
		for col := range b[row] {
			if !b[row][col].Marked {
				sum += b[row][col].Number
			}
		}
	}

	return sum
}

func parseInput(t *testing.T) [][]int {
	f, err := os.Open("example.txt")
	if err != nil {
		t.Fatalf("unable to open file: %s", err)
	}
	defer f.Close()

	var nums []int
	var board Board

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if nums == nil {
			split := strings.Split(scanner.Text(), ",")
			for _, s := range split {
			}
		}

		// scanning a board

		var row []int

		for _, r := range scanner.Text() {
			n, _ := strconv.Atoi(string(r))
			row = append(row, n)
		}

		grid = append(grid, row)
	}

	return grid
}

func TestPartOne(t *testing.T) {
	grid := parseInput(t)
}

func TestPartTwo(t *testing.T) {
}
