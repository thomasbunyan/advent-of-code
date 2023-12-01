package main

import (
	"fmt"
	"strconv"

	"github.com/thomasbunyan/advent-of-code/common"
)

func main() {
	input := common.ReadFast("./internal/2022/day/8/input.txt")

	treeMap := make([][]int, len(input))

	for y, line := range input {
		row := make([]int, len(line))
		for x, tree := range line {
			row[x] = runeToInt(tree)
		}
		treeMap[y] = row
	}

	count := (len(treeMap) * 2) + ((len(treeMap[0]) - 2) * 2)

	for y := 1; y < len(treeMap)-1; y++ {
		for x := 1; x < len(treeMap[y])-1; x++ {
		out:
			for i := -1; i <= 1; i++ {
				for j := -1; j <= 1; j++ {
					if (i == 0 && j == 0) || (i != 0 && j != 0) {
						continue
					}
					if checkVisibility(treeMap, y, x, i, j) {
						count++
						break out
					}
				}
			}
		}
	}

	fmt.Printf("Part 1: %d\n", count)

	maxScore := 0

	for y := 1; y < len(treeMap)-1; y++ {
		for x := 1; x < len(treeMap[y])-1; x++ {
			score := []int{}
			for i := -1; i <= 1; i++ {
				for j := -1; j <= 1; j++ {
					if (i == 0 && j == 0) || (i != 0 && j != 0) {
						continue
					}
					score = append(score, checkScore(treeMap, y, x, i, j))
				}
			}
			total := 1
			for _, s := range score {
				total *= s
			}
			if total > maxScore {
				maxScore = total
			}
		}
	}

	fmt.Printf("Part 2: %d\n", maxScore)
}

func checkVisibility(inMap [][]int, y int, x int, dirY int, dirX int) bool {
	treeSize := inMap[y][x]
	y += dirY
	x += dirX
	for y < len(inMap) && y >= 0 && x < len(inMap[0]) && x >= 0 {
		if treeSize <= inMap[y][x] {
			return false
		}
		y += dirY
		x += dirX
	}
	return true
}

func checkScore(inMap [][]int, y int, x int, dirY int, dirX int) int {
	treeSize := inMap[y][x]
	y += dirY
	x += dirX
	score := 0
	for y < len(inMap) && y >= 0 && x < len(inMap[0]) && x >= 0 {
		if treeSize <= inMap[y][x] {
			return score + 1
		}
		score++
		y += dirY
		x += dirX
	}
	return score
}

func printMap(inMap [][]int) {
	for y := 0; y < len(inMap); y++ {
		for x := 0; x < len(inMap[y]); x++ {
			fmt.Print(inMap[y][x])
		}
		fmt.Print("\n")
	}
}

func runeToInt(in rune) int {
	i, err := strconv.Atoi(fmt.Sprintf("%c", in))
	if err != nil {
		panic(err)
	}
	return i
}
