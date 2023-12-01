package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/thomasbunyan/advent-of-code/common"
)

func main() {
	input := common.ReadFast("./internal/2022/day/10/input.txt")

	cycle := 0
	x := 1

	total := 0

	for _, line := range input {
		command := strings.Split(line, " ")

		pending := 0
		if command[0] == "noop" {
			pending++
			cycle, total = runCycle(pending, cycle, total, x)
		} else if command[0] == "addx" {
			pending += 2
			cycle, total = runCycle(pending, cycle, total, x)
			x += strToInt(command[1])
		}

	}

	fmt.Printf("Total: %d\n", total)
}

func runCycle(pending, cycle, total, x int) (int, int) {
	for ; pending > 0; pending-- {
		cycle++

		sprite := strings.Repeat(".", 40)
		if x >= 0 && x < len(sprite) {
			sprite = sprite[:x] + string("###") + sprite[x+1:]
		} else if x == -1 {
			sprite = sprite[:0] + string("##") + sprite[2:]
		} else if x == 41 {
			sprite = sprite[:len(sprite)-1] + string("##") + sprite[len(sprite)-1:]
		}

		crtRow := ""

		if sprite[cycle%40] == '#' {
			crtRow += "#"
		} else {
			crtRow += "."
		}

		fmt.Print(crtRow)

		if (cycle%40 == 0) && cycle != 0 {
			crtRow = ""
			fmt.Println()
		}

		if (cycle-20)%40 == 0 {
			total += (cycle) * x
		}
	}
	return cycle, total
}

func strToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return i
}
