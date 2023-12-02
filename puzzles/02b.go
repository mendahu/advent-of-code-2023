package puzzles

import (
	"advent/utils"
	"fmt"
	"strconv"
	"strings"
)

func getGamePower(data []string) int {
	val := 0

	for _, game := range data {
		_, results := getGameData(game)

		red := 0
		green := 0
		blue := 0

		for _, result := range strings.Split(results, ";") {
			for _, pull := range strings.Split(result, ",") {
				var count int
				var colour string
				fmt.Sscanf(pull, "%d %s", &count, &colour)

				if (colour == "red" && count > red) {
					red = count
				}

				if (colour == "green" && count > green) {
					green = count
				}

				if (colour == "blue" && count > blue) {
					blue = count
				}
			}
		}

		val = val + (red * green * blue)
	}

	return val
}	

func Puzzle02b() string {
	data := utils.FileReader("data/02.txt")

	val := getGamePower(data)

	return strconv.Itoa(val)
}