package day6

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestPartOne(t *testing.T) {
	fish := getFish(t)
	for i := 0; i < 80; i++ {
		fish = runDay(fish)
	}

	t.Logf("fish: %v", len(fish))
}

func getFish(t *testing.T) []int {
	f, err := os.Open("input.txt")
	if err != nil {
		t.Fatalf("unable to open file: %s", err)
	}
	defer f.Close()

	var fish []int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), ",")
		for _, s := range split {
			n, _ := strconv.Atoi(s)
			fish = append(fish, n)
		}

		// only one line
		break
	}

	return fish
}

func runDay(fish []int) []int {
	n := len(fish)
	for i := 0; i < n; i++ {
		switch fish[i] {
		case 0:
			fish[i] = 6
			fish = append(fish, 8)
		default:
			fish[i]--
		}
	}

	return fish
}

func TestPartTwo(t *testing.T) {
	fish := getFish(t)

	// bucket[i] = the number of fish that are i in their cycle
	bucket := make([]int, 9)

	// place inital fish into initial buckets
	for _, fish := range fish {
		bucket[fish]++
	}

	// run simulation
	for i := 0; i < 256; i++ {
		bucket[0], bucket[1], bucket[2], bucket[3], bucket[4], bucket[5], bucket[6], bucket[7], bucket[8] = bucket[1], bucket[2], bucket[3], bucket[4], bucket[5], bucket[6], bucket[7]+bucket[0], bucket[8], bucket[0]
	}

	// count fish
	sum := 0
	for _, count := range bucket {
		sum += count
	}

	t.Logf("Fish: %v", sum)
}
