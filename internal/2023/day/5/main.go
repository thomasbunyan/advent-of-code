package main

import (
	"fmt"
	"strings"

	"github.com/thomasbunyan/advent-of-code/common"
)

func main() {
	lines := common.ReadFast("./internal/2023/day/5/input.txt")

	almanacs := []func(in int) int{}
	almanacsReverse := []func(in int) int{}

	seeds := common.StrToIntSlice(strings.Fields(strings.Split(lines[0], ":")[1]))

	for i := 1; i < len(lines); i++ {
		if lines[i] != "" && lines[i][len(lines[i])-1] == ':' {
			var values = [][]int{}
			i++
			for i < len(lines) && lines[i] != "" {
				values = append(values, common.StrToIntSlice(strings.Fields(lines[i])))
				i++
			}
			almanacs = append(almanacs, func(in int) int {
				for _, value := range values {
					if in >= value[1] && in < value[1]+value[2] {
						return in + (value[0] - value[1])
					}
				}
				return in
			})
			almanacsReverse = append(almanacsReverse, func(in int) int {
				for _, value := range values {
					if in-(value[0]-value[1]) >= value[1] && in-(value[0]-value[1]) < value[1]+value[2] {
						return in - (value[0] - value[1])
					}
				}
				return in
			})
		}
	}

	lowest := -1

	for i := 0; i < len(seeds); i++ {
		input := seeds[i]
		for _, convert := range almanacs {
			input = convert(input)
		}
		if lowest == -1 || input < lowest {
			lowest = input
		}
	}

	fmt.Printf("Ans 1: %d\n", lowest)

	lowest = -1

	for i := 1; ; i++ {
		score := i
		for j := len(almanacsReverse) - 1; j >= 0; j-- {
			score = almanacsReverse[j](score)
		}
		if validSeed(seeds, score) {
			lowest = i
			break
		}
	}

	fmt.Printf("Ans 2: %d\n", lowest)
}

func validSeed(seeds []int, score int) bool {
	for i := 0; i < len(seeds); i += 2 {
		if score >= seeds[i] && score < seeds[i]+seeds[i+1] {
			return true
		}
	}
	return false
}
