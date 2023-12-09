package main

import (
	"fmt"
	"strings"

	"github.com/thomasbunyan/advent-of-code/common"
)

func main() {
	lines := common.ReadFast("./internal/2023/day/9/input.txt")

	nextTotal := 0
	prevTotal := 0

	for _, line := range lines {
		values := common.StrToIntSlice(strings.Fields(line))
		next, prev := getHistory(values)
		nextTotal += next
		prevTotal += prev
	}

	fmt.Printf("Ans 1: %d\n", nextTotal)
	fmt.Printf("Ans 2: %d\n", prevTotal)
}

func getHistory(values []int) (int, int) {
	if allEqual(values) {
		return values[0], values[0]
	}

	deltas := make([]int, len(values)-1)
	for i := 1; i < len(values); i++ {
		deltas[i-1] = values[i] - values[i-1]
	}

	next, prev := getHistory(deltas)
	return values[len(values)-1] + next, values[0] - prev
}

func allEqual(s []int) bool {
	for _, delta := range s[1:] {
		if s[0] != delta {
			return false
		}
	}
	return true
}
