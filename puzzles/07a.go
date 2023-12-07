package puzzles

import (
	"advent/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type CardHand int

const (
	FiveOfAKind     CardHand = 7
	FourOfAKind     CardHand = 6
	FullHouse       CardHand = 5
	ThreeOfAKind    CardHand = 4
	TwoPairs        CardHand = 3
	OnePair         CardHand = 2
	HighCard        CardHand = 1
)

type Hand struct {
	Cards []string
	Bid int
	Hand CardHand
}

func (h *Hand) CalculateHand(useJoker bool) {
	cardCounts := map[string]int{}

	jokerCount := 0
	for _, card := range h.Cards {
		if (card == "J") {
			jokerCount++
		}
		cardCounts[card]++
	}

	// Check for five of a kind
	if (len(cardCounts) == 1) {
		h.Hand = FiveOfAKind
		return
	}

	// Check for four of a kind // Full House
	if (len(cardCounts) == 2) {
		if (useJoker && jokerCount > 0) {
			h.Hand = FiveOfAKind
			return
		}

		if (cardCounts[h.Cards[0]] == 4 || cardCounts[h.Cards[0]] == 1) {
			h.Hand = FourOfAKind
			return
		} else {
			h.Hand = FullHouse
			return
		}
	}

	// Check for three of a kind // Two Pairs
	if (len(cardCounts) == 3) {
		if (cardCounts[h.Cards[0]] == 2 || cardCounts[h.Cards[1]] == 2 || cardCounts[h.Cards[2]] == 2) {
			if (useJoker && jokerCount > 0) {
				if (jokerCount > 1) {
					h.Hand = FourOfAKind
					return
				}

				h.Hand = FullHouse
				return
			}
			h.Hand = TwoPairs
			return
		} else {
			if (useJoker && jokerCount > 0) {
				h.Hand = FourOfAKind
				return
			}
			h.Hand = ThreeOfAKind
			return
		}
	}

	// Check for one pair // High Card
	if (len(cardCounts) == 4) {
		if (useJoker && jokerCount > 0) {
			h.Hand = ThreeOfAKind
			return
		}
		h.Hand = OnePair
		return
	} else {
		if (useJoker && jokerCount > 0) {
			h.Hand = OnePair
			return
		}
		h.Hand = HighCard
		return
	}
}

func getHandsData(data []string, useJoker bool) []Hand {
	hands := make([]Hand, 0)

	for _, line := range data {
		hand := Hand{}

		var cards string
		var bid int
		fmt.Sscanf(line , "%s %d", &cards, &bid)

		hand.Cards = strings.Split(cards, "")
		hand.Bid = bid

		hand.CalculateHand(useJoker)

		hands = append(hands, hand)
	}

	return hands
}

func generateHandsSorter(hands []Hand, useJoker bool) func (i, j int) bool {
	return func (i, j int) bool {
		cards := map[string]int{
			"2": 2,
			"3": 3,
			"4": 4,
			"5": 5,
			"6": 6,
			"7": 7,
			"8": 8,
			"9": 9,
			"T": 10,
			"Q": 12,
			"K": 13,
			"A": 14,
		}

		if (useJoker) {
			cards["J"] = 1
		} else {
			cards["J"] = 11
		}

		// card is the same, do secondary ordering
		if (hands[i].Hand == hands[j].Hand) {

			for k := 0; k < len(hands[i].Cards); k++ {
				if (cards[hands[i].Cards[k]] == cards[hands[j].Cards[k]]) {
					continue
				}
				
				return cards[hands[i].Cards[k]] > cards[hands[j].Cards[k]]
			}
		}

		// Otherwise return best hand
		return hands[i].Hand > hands[j].Hand
	}
}

func Puzzle07a() string {
	data := utils.FileReader("data/07_test_2.txt")
	hands := getHandsData(data, false)

	score := 0

	useJoker := false

	sort.Slice(hands, generateHandsSorter(hands, useJoker))

	for i, rank := 0, len(hands); i < len(hands); i, rank = i+1, rank-1 {
		winnings := hands[i ].Bid * rank
		fmt.Println(hands[i], rank, winnings)
		score = score + winnings
	}

	return strconv.Itoa(score)
}