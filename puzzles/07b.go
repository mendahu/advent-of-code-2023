package puzzles

import (
	"advent/utils"
	"fmt"
	"sort"
	"strconv"
)

func Puzzle07b() string {
	data := utils.FileReader("data/07.txt")
	hands := getHandsData(data, true)

	score := 0

	useJoker := true

	sort.Slice(hands, generateHandsSorter(hands, useJoker))

	for i, rank := 0, len(hands); i < len(hands); i, rank = i+1, rank-1 {
		winnings := hands[i].Bid * rank
		fmt.Println(hands[i], rank, winnings)
		score = score + winnings
	}

	return strconv.Itoa(score)
}

// not 250563526