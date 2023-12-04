package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/thomasbunyan/advent-of-code/common"
)

func main() {
	lines := common.ReadFast("./internal/2023/day/4/input.txt")

	score := 0

	copies := make(map[int]int)

	for _, line := range lines {
		card := strings.Split(line, ":")
		numbers := strings.Split(card[1], "|")

		cardNo := common.StrToInt(strings.Fields(card[0])[1])
		copies[cardNo]++

		winning := strings.Fields(strings.TrimSpace(numbers[0]))
		ourHand := strings.Fields(strings.TrimSpace(numbers[1]))

		numberSet := make(map[string]bool)

		for _, no := range winning {
			numberSet[no] = true
		}

		handScore := 0

		for _, no := range ourHand {
			if numberSet[no] {
				handScore++
			}
		}

		if handScore != 0 {
			score += int(math.Pow(2, float64(handScore-1)))

			for i := 1; i <= handScore; i++ {
				copies[cardNo+i] = (copies[cardNo+i] + copies[cardNo])
			}
		}
	}

	scoreB := 0
	for _, count := range copies {
		scoreB += count
	}

	fmt.Printf("Ans 1: %d\n", score)
	fmt.Printf("Ans 2: %d\n", scoreB)
}
