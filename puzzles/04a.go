package puzzles

import (
	"advent/utils"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func getLotteryNumbers(data string) ([]string, []string) {
	components := strings.Split(data, ": ")
	sets := strings.Split(components[1], " | ")

	re := regexp.MustCompile(`([0-9]+)`)

	winners := re.FindAllString(sets[0], -1)
	picks := re.FindAllString(sets[1], -1)

	return winners, picks
}

func Puzzle04a() string {
	data := utils.FileReader("data/04.txt")

	var val int = 0

	for _, line := range data {
		winners, picks := getLotteryNumbers(line)

		bits := 0

		for _, picks := range picks {
			for _, winner := range winners {
				if picks == winner {
					bits += 1
				}
			}
		}

		points := math.Pow(2, float64(bits - 1))
		fmt.Println(bits, points)
		val += int(points)
	}


	return strconv.Itoa(int(val))
}