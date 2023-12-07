package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/thomasbunyan/advent-of-code/common"
)

const HAND_SIZE = 5

var CARD_STRENGTHS = map[byte]int{
	'2': 0,
	'3': 1,
	'4': 2,
	'5': 3,
	'6': 4,
	'7': 5,
	'8': 6,
	'9': 7,
	'T': 8,
	'J': 9,
	'Q': 10,
	'K': 11,
	'A': 12,
}

func main() {
	lines := common.ReadFast("./internal/2023/day/7/input.txt")

	for part := 1; part <= 2; part++ {
		joker := false
		if part == 2 {
			CARD_STRENGTHS['J'] = -1
			joker = true
		}

		sort.Slice(lines, func(i, j int) bool {
			handI := strings.Fields(lines[i])[0]
			handJ := strings.Fields(lines[j])[0]

			handIScore := getHandScore(handI, joker)
			handJScore := getHandScore(handJ, joker)

			if handIScore == handJScore {
				for card := 0; card < HAND_SIZE; card++ {
					if handI[card] == handJ[card] {
						continue
					} else {
						return CARD_STRENGTHS[handI[card]] > CARD_STRENGTHS[handJ[card]]
					}
				}
				return false
			} else {
				return handIScore > handJScore
			}
		})

		score := 0

		for rank, line := range lines {
			score += (len(lines) - rank) * common.StrToInt(strings.Fields(line)[1])
		}

		fmt.Printf("Ans %d: %d\n", part, score)
	}
}

func getHandScore(hand string, joker bool) int {
	cards := map[rune]int{}
	jokerCount := 0
	for _, card := range hand {
		if joker && card == 'J' {
			jokerCount++
		} else {
			cards[card]++
		}
	}
	if len(cards) == 0 || len(cards) == 1 {
		// Five of a kind
		return 6
	} else if len(cards) == 2 {
		for _, count := range cards {
			if count == 4-jokerCount {
				// Four of a kind
				return 5
			}
		}
		// Full house
		return 4
	} else if len(cards) == 3 {
		for _, count := range cards {
			if count == 3-jokerCount {
				// Three of a kind
				return 3
			}
		}
		// Two pair
		return 2
	} else if len(cards) == 4 {
		// One pair
		return 1
	} else if len(cards) == 5 {
		// High card
		return 0
	} else {
		panic("bad hand")
	}
}
