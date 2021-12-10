package day10

import (
	"bufio"
	"os"
	"sort"
	"testing"
)

type Stack []rune

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(str rune) {
	*s = append(*s, str)
}

func (s *Stack) Pop() (rune, bool) {
	if s.IsEmpty() {
		return 0, false
	}

	idx := len(*s) - 1
	element := (*s)[idx]
	*s = (*s)[:idx]
	return element, true
}

func isOpenSymbol(r rune) bool {
	return r == '(' || r == '[' || r == '{' || r == '<'
}

func symbolsMatch(open, close rune) bool {
	switch open {
	case '(':
		return close == ')'
	case '[':
		return close == ']'
	case '{':
		return close == '}'
	case '<':
		return close == '>'
	}

	return false
}

func parseInput(t *testing.T) []string {
	f, err := os.Open("input.txt")
	if err != nil {
		t.Fatalf("unable to open file: %s", err)
	}
	defer f.Close()

	var lines []string

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

var points = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

func TestPartOne(t *testing.T) {
	lines := parseInput(t)
	score := 0

	for _, line := range lines {
		valid, expected := isValid(line)
		if !valid {
			score += points[expected]
		}
	}

	t.Logf("Score: %v", score)
}

// isValid returns if the line was valid, and, if it's not valid, the first invalid rune.
func isValid(line string) (bool, rune) {
	var stack Stack

	for _, r := range line {
		if isOpenSymbol(r) {
			stack.Push(r)
			continue
		}

		open, ok := stack.Pop()
		if !ok {
			// no symbols left to pop (weren't expecting anything?)
			return false, 0
		}

		if !symbolsMatch(open, r) {
			return false, r
		}
	}

	return true, 0
}

func TestPartTwo(t *testing.T) {
	lines := parseInput(t)

	var incomplete []string

	for _, line := range lines {
		valid, _ := isValid(line)
		if valid {
			incomplete = append(incomplete, line)
		}
	}

	var scores []int

	for _, line := range incomplete {
		completion := getCompletionString(line)
		score := completionScore(completion)
		scores = append(scores, score)
	}

	sort.Ints(scores)
	t.Logf("Mid Score: %v", scores[len(scores)/2])
}

var opposite = map[rune]rune{
	'(': ')',
	'[': ']',
	'{': '}',
	'<': '>',
}

func getCompletionString(line string) []rune {
	var stack Stack

	for _, r := range line {
		if isOpenSymbol(r) {
			stack.Push(r)
			continue
		}

		open, ok := stack.Pop()
		if !ok {
			return nil // shouldn't happen
		}

		if !symbolsMatch(open, r) {
			return nil // also shouldn't happen
		}
	}

	var completion []rune
	for i := len(stack) - 1; i >= 0; i-- {
		completion = append(completion, opposite[stack[i]])
	}

	return completion
}

func completionScore(completion []rune) int {
	score := 0

	for _, r := range completion {
		score *= 5

		switch r {
		case ')':
			score += 1
		case ']':
			score += 2
		case '}':
			score += 3
		case '>':
			score += 4
		}
	}

	return score
}
