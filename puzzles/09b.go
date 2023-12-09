package puzzles

import (
	"advent/utils"
	"strconv"
)

func (s *Sequence) GetFirstVal() int {
	return s.Values[0]
}

func Puzzle09b() string {
	data := utils.FileReader("data/09.txt")
	sequences := getSequences(data)

	sum := 0

	for _, s := range sequences {
		incrementor := 0
		sign := 1

		var recursivelyGetSequences func([]int)
		
		recursivelyGetSequences = func(sequence []int) {
			sequence, isAllZeroes := GetNextSequence(sequence)

			if (isAllZeroes) {
				sign = sign * -1
				sum += s.GetFirstVal() + incrementor 
			} else {
				sign = sign * -1
				incrementor = incrementor + sequence[0] * sign
				recursivelyGetSequences(sequence)
			}
		}

		recursivelyGetSequences(s.Values)
	}

	return strconv.Itoa(sum)
}