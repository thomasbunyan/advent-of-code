package main

import (
	"fmt"
	"strings"

	"github.com/thomasbunyan/advent-of-code/common"
)

func main() {
	lines := common.ReadFast("./internal/2023/day/2/input.txt")

	partOne(lines)
	partTwo(lines)
}

func partOne(lines []string) {
	MAX_RED, MAX_GREEN, MAX_BLUE := 12, 13, 14

	score := 0

	for _, line := range lines {
		gamePossible := true

		gameSplit := strings.Split(line, ":")
		rounds := strings.Split(gameSplit[1], ";")

	out:
		for _, round := range rounds {
			cubes := strings.Split(round, ",")
			for _, cube := range cubes {
				colorCount := strings.Split(strings.Trim(cube, " "), " ")
				if colorCount[1] == "red" {
					if common.StrToInt(colorCount[0]) > MAX_RED {
						gamePossible = false
						break out
					}
				}
				if colorCount[1] == "green" {
					if common.StrToInt(colorCount[0]) > MAX_GREEN {
						gamePossible = false
						break out
					}
				}
				if colorCount[1] == "blue" {
					if common.StrToInt(colorCount[0]) > MAX_BLUE {
						gamePossible = false
						break out
					}
				}
			}
		}

		if gamePossible {
			score += common.StrToInt(strings.Split(gameSplit[0], " ")[1])
		}
	}

	fmt.Printf("Part 1: %d\n", score)
}

func partTwo(lines []string) {
	score := 0

	for _, line := range lines {
		gameSplit := strings.Split(line, ":")
		rounds := strings.Split(gameSplit[1], ";")

		minRed, minGreen, minBlue := -1, -1, -1

		for _, round := range rounds {
			cubes := strings.Split(round, ",")
			for _, cube := range cubes {
				colorCount := strings.Split(strings.Trim(cube, " "), " ")
				if colorCount[1] == "red" {
					if common.StrToInt(colorCount[0]) > minRed {
						minRed = common.StrToInt(colorCount[0])
					}
				}
				if colorCount[1] == "green" {
					if common.StrToInt(colorCount[0]) > minGreen {
						minGreen = common.StrToInt(colorCount[0])
					}
				}
				if colorCount[1] == "blue" {
					if common.StrToInt(colorCount[0]) > minBlue {
						minBlue = common.StrToInt(colorCount[0])
					}
				}
			}
		}

		score += minRed * minGreen * minBlue
	}

	fmt.Printf("Part 2: %d\n", score)
}
