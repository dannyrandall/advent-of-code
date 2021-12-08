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

	var nums []int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatalf("unable to parse number: %s", err)
		}

		nums = append(nums, n)
	}

	increased := 0

	curSum := nums[0] + nums[1] + nums[2]
	for i := 3; i < len(nums); i++ {
		nextSum := curSum - nums[i-3] + nums[i]
		if nextSum > curSum {
			increased++
		}

		curSum = nextSum
	}

	log.Printf("increased: %v", increased)
}
