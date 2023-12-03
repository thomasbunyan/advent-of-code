package main

import (
	"fmt"
	"strconv"

	"github.com/thomasbunyan/advent-of-code/common"
)

func main() {
	lines := common.ReadFast("./internal/2023/day/3/input.txt")

	partOne(lines)
	partTwo(lines)
}

func partOne(lines []string) {
	score := 0

	for j, line := range lines {
		for i := 0; i < len(line); i++ {
			var strInt string
			isPart := false
			for i < len(line) && isInt(string(line[i])) {
				strInt += string(line[i])
				if !isPart {
					isPart = checkPart(lines, j, i)
				}
				i++
			}
			if strInt != "" && isPart {
				score += common.StrToInt(strInt)
			}
		}
	}

	fmt.Printf("Ans: %d\n", score)
}

func partTwo(lines []string) {
	gears := make(map[string][]int)

	for j, line := range lines {
		for i := 0; i < len(line); i++ {
			var strInt string
			gearSet := make(map[string]bool)
			for i < len(line) && isInt(string(line[i])) {
				strInt += string(line[i])
				for _, gear := range checkGears(lines, j, i) {
					gearSet[gear] = true
				}
				i++
			}
			if strInt != "" && len(gearSet) > 0 {
				for gear := range gearSet {
					gears[gear] = append(gears[gear], common.StrToInt(strInt))
				}
			}
		}
	}

	score := 0

	for _, part := range gears {
		if len(part) == 2 {
			score += part[0] * part[1]
		}
	}

	fmt.Printf("Ans: %d\n", score)
}

func checkPart(schematic []string, y, x int) bool {
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i != 0 || j != 0 {
				if y+i >= 0 && y+i < len(schematic) && x+j >= 0 && x+j < len(schematic[y]) {
					position := string(schematic[y+i][x+j])
					if position != "." && !isInt(position) {
						return true
					}
				}
			}
		}
	}
	return false
}

func checkGears(schematic []string, y, x int) []string {
	var gears []string
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i != 0 || j != 0 {
				if y+i >= 0 && y+i < len(schematic) && x+j >= 0 && x+j < len(schematic[y]) {
					if string(schematic[y+i][x+j]) == "*" {
						gears = append(gears, fmt.Sprintf("%d,%d", y+i, x+j))
					}
				}
			}
		}
	}
	return gears
}

func isInt(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}
