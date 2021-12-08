package day3

import (
	"bufio"
	"os"
	"strconv"
	"testing"
)

func TestPartOne(t *testing.T) {
	f, err := os.Open("input.txt")
	if err != nil {
		t.Fatalf("unable to open file: %s", err)
	}
	defer f.Close()

	ones := make([]int, 12)
	zeros := make([]int, 12)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		for i, r := range scanner.Text() {
			switch r {
			case '0':
				zeros[i]++
			case '1':
				ones[i]++
			default:
				t.Fatalf("wtf")
			}
		}
	}

	gamma := ""
	epislon := ""

	for i := range zeros {
		if zeros[i] > ones[i] {
			gamma += "0"
			epislon += "1"
			continue
		}

		gamma += "1"
		epislon += "0"
	}

	g, _ := strconv.ParseInt(gamma, 2, 64)
	e, _ := strconv.ParseInt(epislon, 2, 64)

	t.Logf("g*e: %v", g*e)
}

func TestPartTwo(t *testing.T) {
	f, err := os.Open("input.txt")
	if err != nil {
		t.Fatalf("unable to open file: %s", err)
	}
	defer f.Close()

	var nums []string

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		nums = append(nums, scanner.Text())
	}

	// determine oxygen generator rating
	oxy := make([]string, len(nums))
	copy(oxy, nums)

	oxyRating := ""
	for i := 0; i < len(nums[0]); i++ {
		oxy = filterMostCommon(oxy, i)
		if len(oxy) == 1 {
			oxyRating = oxy[0]
			break
		}
	}

	// determine c02 scrubber rating
	scrubber := make([]string, len(nums))
	copy(scrubber, nums)

	scrubberRating := ""
	for i := 0; i < len(nums[0]); i++ {
		scrubber = filterLeastCommon(scrubber, i)
		if len(scrubber) == 1 {
			scrubberRating = scrubber[0]
			break
		}
	}

	a, _ := strconv.ParseInt(oxyRating, 2, 64)
	b, _ := strconv.ParseInt(scrubberRating, 2, 64)

	t.Logf("a * b: %v * %v: %v", a, b, a*b)
}

func filterMostCommon(nums []string, pos int) []string {
	var out []string
	common := mostCommon(nums, pos)

	for _, num := range nums {
		if rune(num[pos]) == common {
			out = append(out, num)
		}
	}

	return out
}

func filterLeastCommon(nums []string, pos int) []string {
	var out []string
	common := leastCommon(nums, pos)

	for _, num := range nums {
		if rune(num[pos]) == common {
			out = append(out, num)
		}
	}

	return out
}

func mostCommon(nums []string, pos int) rune {
	ones := 0
	zeros := 0

	for _, num := range nums {
		if num[pos] == '0' {
			zeros++
			continue
		}

		ones++
	}

	switch {
	case ones == zeros:
		return '1'
	case ones > zeros:
		return '1'
	default:
		return '0'
	}
}

func leastCommon(nums []string, pos int) rune {
	ones := 0
	zeros := 0

	for _, num := range nums {
		if num[pos] == '0' {
			zeros++
			continue
		}

		ones++
	}

	switch {
	case ones == zeros:
		return '0'
	case ones > zeros:
		return '0'
	default:
		return '1'
	}
}
