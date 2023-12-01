package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/thomasbunyan/advent-of-code/common"
)

func main() {
	input := common.ReadFast("./internal/2022/day/9/input.txt")

	const CHAIN_SIZE = 9

	headY, headX := 0, 0
	tail := [CHAIN_SIZE][2]int{}

	tailLog := make(map[string]bool)

	for _, line := range input {
		command := strings.Split(line, " ")
		momentumY, momentumX := getMomentum(command[0])
		for i := 0; i < strToInt(command[1]); i++ {
			headY += momentumY
			headX += momentumX

			for tailId, node := range tail {
				var prevY, prevX int
				if tailId == 0 {
					prevY, prevX = headY, headX
				} else {
					prevY, prevX = tail[tailId-1][0], tail[tailId-1][1]
				}

				tail[tailId] = makeTuple(moveTail(prevY, prevX, node[0], node[1]))
				tailLog[fmt.Sprintf("%d,%d", tail[len(tail)-1][0], tail[len(tail)-1][1])] = true
			}
		}
	}

	fmt.Printf("Ans: %d\n", len(tailLog))
}

func moveTail(headY, headX, tailY, tailX int) (int, int) {
	diffY := diff(headY, tailY)
	diffX := diff(headX, tailX)

	moveY, moveX := 0, 0

	if diffY == 2 && diffX == 0 {
		if moveY = 1; headY < tailY {
			moveY = -1
		}
	} else if diffX == 2 && diffY == 0 {
		if moveX = 1; headX < tailX {
			moveX = -1
		}
	} else if diffY > 1 || diffX > 1 {
		if moveY = 1; headY < tailY {
			moveY = -1
		}
		if moveX = 1; headX < tailX {
			moveX = -1
		}
	}

	return tailY + (moveY), tailX + (moveX)
}

func getMomentum(dir string) (int, int) {
	switch dir {
	case "R":
		return 0, 1
	case "U":
		return -1, 0
	case "L":
		return 0, -1
	case "D":
		return 1, 0
	default:
		panic("Invalid direction")
	}
}

func makeTuple(a, b int) [2]int {
	return [2]int{a, b}
}

func diff(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}

func strToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return i
}
