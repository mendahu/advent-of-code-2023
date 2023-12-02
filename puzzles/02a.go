package puzzles

import (
	"advent/utils"
	"fmt"
	"strconv"
	"strings"
)

func getGameData(game string) (id int, results string) {
	parts := strings.Split(game, ": ")
	
	fmt.Sscanf(parts[0], "Game %d", &id)
	results = parts[1]
	return id, results
}

func testGame(data []string, red int, green int, blue int) int {
	val := 0
	
	for _, game := range data {
		gameNum, results := getGameData(game)

		isPossible := true

		exit:
		for _, result := range strings.Split(results, ";") {
			fmt.Println(result)
			for _, pull := range strings.Split(result, ",") {
				var count int
				var colour string
				fmt.Sscanf(pull, "%d %s", &count, &colour)

				if (colour == "red" && count > red) {
					isPossible = false
					break exit;
				}

				if (colour == "green" && count > green) {
					isPossible = false
					break exit;
				}

				if (colour == "blue" && count > blue) {
					isPossible = false
					break exit;
				}
			}
		}

		if (isPossible) {
			val = val + gameNum
		}
	}

	return val;
}

func Puzzle02a() string {
	data := utils.FileReader("data/02.txt")

	val := testGame(data, 12, 13, 14)

	return strconv.Itoa(val)
}