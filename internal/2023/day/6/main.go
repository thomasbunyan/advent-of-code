package main

import (
	"fmt"
	"strings"

	"github.com/thomasbunyan/advent-of-code/common"
)

func main() {
	lines := common.ReadFast("./internal/2023/day/6/input.txt")

	times := common.StrToIntSlice(strings.Fields(lines[0])[1:])
	distances := common.StrToIntSlice(strings.Fields(lines[1])[1:])

	score := 1

	for i := 0; i < len(times); i++ {
		score *= getPermutations(times[i], distances[i])
	}

	fmt.Printf("Ans 1: %d\n", score)

	timeBig := common.StrToInt(strings.Join(strings.Fields(lines[0])[1:], ""))
	distanceBig := common.StrToInt(strings.Join(strings.Fields(lines[1])[1:], ""))

	fmt.Printf("Ans 2: %d\n", getPermutations(timeBig, distanceBig))
}

func getPermutations(time, distance int) int {
	min := 0
	max := 0
	for j := 0; j < time; j++ {
		if (j * (time - j)) > distance {
			min = j
			break
		}
	}
	for j := time; j >= 0; j-- {
		if (j * (time - j)) > distance {
			max = j
			break
		}
	}
	return ((max - min) + 1)
}
