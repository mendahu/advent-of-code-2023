package puzzles

import (
	"advent/utils"
	"strconv"
	"strings"
)

func Puzzle01b() string {
	data := utils.FileReader("data/01.txt")

	result := 0

	stringMap := map[string]int{
		"one": 1,
		"two": 2,
		"three": 3,
		"four": 4,
		"five": 5,
		"six": 6,
		"seven": 7,
		"eight": 8,
		"nine": 9,
		"1": 1,
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
	}
	
	for _, line := range data {
		var vals []int

		for i := range line {
			for str, val := range stringMap {
				if strings.HasPrefix(line[i:], str) {
					vals = append(vals, val)
				}
			}
		}
		
		result = result + (vals[0] * 10) + vals[len(vals)-1]
	}

	return strconv.Itoa(result)
}