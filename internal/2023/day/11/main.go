package main

import (
	"fmt"
	"math"

	"github.com/thomasbunyan/advent-of-code/common"
)

const EXPANSION_FACTOR = 1_000_000

func main() {
	lines := common.ReadFast("./internal/2023/day/11/input.txt")

	galaxy := make([][]rune, len(lines))

	for y, line := range lines {
		galaxy[y] = make([]rune, len(line))
		for x, space := range line {
			galaxy[y][x] = space
		}
	}

	emptyCols, emptyRows := getEmpty(galaxy)

	galaxyCoords := [][]int{}
	for y := range galaxy {
		expansionY := 0
		for _, emptyRow := range emptyRows {
			if y > emptyRow {
				expansionY += EXPANSION_FACTOR - 1
			}
		}
		for x := range galaxy[y] {
			expansionX := 0
			for _, emptyCol := range emptyCols {
				if x > emptyCol {
					expansionX += EXPANSION_FACTOR - 1
				}
			}
			if galaxy[y][x] == '#' {
				galaxyCoords = append(galaxyCoords, []int{y + expansionY, x + expansionX})
			}
		}
	}

	score := 0

	for i := 0; i < len(galaxyCoords); i++ {
		for j := i + 1; j < len(galaxyCoords); j++ {
			score += getShortestPath(galaxyCoords[i], galaxyCoords[j])
		}
	}

	fmt.Printf("Ans: %d\n", score)
}

func getShortestPath(a, b []int) int {
	return int(math.Abs(float64(a[0]-b[0])) + math.Abs(float64(a[1]-b[1])))
}

func getEmpty(galaxy [][]rune) ([]int, []int) {
	emptyCols, emptyRows := []int{}, []int{}

	for y := 0; y < len(galaxy); y++ {
		noGal := true
		for x := 0; x < len(galaxy[0]); x++ {
			if galaxy[y][x] == '#' {
				noGal = false
				break
			}
		}
		if noGal {
			emptyRows = append(emptyRows, y)
		}
	}

	for x := 0; x < len(galaxy[0]); x++ {
		noGal := true
		for y := 0; y < len(galaxy); y++ {
			if galaxy[y][x] == '#' {
				noGal = false
				break
			}
		}
		if noGal {
			emptyCols = append(emptyCols, x)
		}
	}

	return emptyCols, emptyRows
}
