package day8

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

const (
	lineFormat = "%s %s %s %s %s %s %s %s %s %s | %s %s %s %s"
)

func TestPartOne(t *testing.T) {
	f, err := os.Open("input.txt")
	if err != nil {
		t.Fatalf("unable to open file: %s", err)
	}
	defer f.Close()

	unique := make([]string, 10)
	output := make([]string, 4)

	count := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Sscanf(scanner.Text(), lineFormat, &unique[0], &unique[1], &unique[2], &unique[3], &unique[4], &unique[5], &unique[6], &unique[7], &unique[8], &unique[9], &output[0], &output[1], &output[2], &output[3])

		for _, segment := range output {
			switch len(segment) {
			case 2: // is a 1
				count++
			case 4: // is a 4
				count++
			case 3: // is a 7
				count++
			case 7: // is a 8
				count++
			}
		}
	}

	t.Logf("count: %v", count)
}

func TestPartTwo(t *testing.T) {
	f, err := os.Open("input.txt")
	if err != nil {
		t.Fatalf("unable to open file: %s", err)
	}
	defer f.Close()

	unique := make([]string, 10)
	output := make([]string, 4)

	sum := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Sscanf(scanner.Text(), lineFormat, &unique[0], &unique[1], &unique[2], &unique[3], &unique[4], &unique[5], &unique[6], &unique[7], &unique[8], &unique[9], &output[0], &output[1], &output[2], &output[3])

		// digit[<signal>] = the digit <signal> maps to
		digit := make(map[string]int, 10)

		// pattern[i] = the pattern digit i maps to
		pattern := make([]string, 10)

		for _, signal := range unique {
			switch len(signal) {
			case 2: // is a 1
				digit[signal] = 1
				pattern[1] = signal
			case 4: // is a 4
				digit[signal] = 4
				pattern[4] = signal
			case 3: // is a 7
				digit[signal] = 7
				pattern[3] = signal
			case 7: // is a 8
				digit[signal] = 8
				pattern[7] = signal
			}
		}

		// determine 2, 3, 5
		for _, signal := range unique {
			switch {
			case len(signal) != 5:
			case countRuneMatches(pattern[1], signal) == 2: // 3 matches both runes in 1
				digit[signal] = 3
				pattern[3] = signal
			case countRuneMatches(pattern[4], signal) == 2: // 2 matches two runes in 4
				digit[signal] = 2
				pattern[2] = signal
			default: // if neither of those, it's 5
				digit[signal] = 5
				pattern[5] = signal
			}
		}

		// determine 0, 6, 9
		for _, signal := range unique {
			switch {
			case len(signal) != 6:
			case countRuneMatches(pattern[4], signal) == 4: // 9 matches four runes in 4
				digit[signal] = 9
				pattern[9] = signal
			case countRuneMatches(pattern[5], signal) == 5: // 6 matches five runes in 5 (and isn't 9)
				digit[signal] = 6
				pattern[6] = signal
			default: // if neither of those, it's 0
				digit[signal] = 0
				pattern[0] = signal
			}
		}

		// calculate output value
		outputNum := (1000 * findDigit(digit, output[0])) + (100 * findDigit(digit, output[1])) + (10 * findDigit(digit, output[2])) + findDigit(digit, output[3])
		sum += outputNum
	}

	t.Logf("sum: %v", sum)
}

func countRuneMatches(a, b string) int {
	shorter := a
	longer := b

	if len(shorter) > len(longer) {
		shorter, longer = longer, shorter
	}

	inShorter := make(map[rune]bool, len(shorter))
	for _, r := range shorter {
		inShorter[r] = true
	}

	count := 0
	for _, r := range longer {
		if inShorter[r] {
			count++
		}
	}

	return count
}

func findDigit(digit map[string]int, signal string) int {
	digitMatch := func(a, b string) bool {
		if len(a) != len(b) {
			return false
		}

		return countRuneMatches(a, b) == len(a)
	}

	for k, v := range digit {
		if digitMatch(k, signal) {
			return v
		}
	}

	return -1
}
