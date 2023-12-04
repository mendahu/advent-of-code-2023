package puzzles

import (
	"advent/utils"
	"strconv"
)

func Puzzle04b() string {
	data := utils.FileReader("data/04.txt")

	count := 0

	copies := make([]int, len(data))
	for i := range copies {
		copies[i] = 1
	}

	for index, line := range data {
		winners, picks := getLotteryNumbers(line)

		wins := 0

		for _, picks := range picks {
			for _, winner := range winners {
				if picks == winner {
					wins += 1
				}
			}
		}
		
		count += copies[index]

		i := index + 1
		for {
			if (wins <= 0) {
				break
			}
			copies[i] += copies[index]
			i++
			wins--
		}
	}


	return strconv.Itoa(count)
}