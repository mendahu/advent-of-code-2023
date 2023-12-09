package puzzles

import (
	"advent/utils"
	"strconv"
	"strings"
)

type Sequence struct {
	Values []int
}

func (s *Sequence) GetLastVal() int {
	return s.Values[len(s.Values) - 1]
}

func GetNextSequence(sequence []int) ([]int, bool) {
	newSequence := make([]int, 0)
	isAllZeros := true
	for i := 1; i < len(sequence); i++ {
		newVal := sequence[i] - sequence[i - 1]
		newSequence = append(newSequence, newVal)
		if newVal != 0 {
			isAllZeros = false
		}
	}

	return newSequence, isAllZeros
}

func getSequences(data []string) []Sequence {
	var sequences []Sequence

	for _, line := range data {
		valuesStrings := strings.Split(line, " ")
		values := make([]int, len(valuesStrings))

		for i, valueString := range valuesStrings {
			value, _ := strconv.Atoi(valueString)
			values[i] = value
		}

		sequence := Sequence{Values: values}
		sequences = append(sequences, sequence)
	}

	return sequences

}

func Puzzle09a() string {
	data := utils.FileReader("data/09.txt")
	sequences := getSequences(data)

	sum := 0

	for _, s := range sequences {
		incrementor := 0

		var recursivelyGetSequences func([]int)
		
		recursivelyGetSequences = func(sequence []int) {
			sequence, isAllZeroes := GetNextSequence(sequence)

			if (isAllZeroes) {

				sum += incrementor + s.GetLastVal()
			} else {
				incrementor = incrementor + sequence[len(sequence) - 1]
				recursivelyGetSequences(sequence)
			}
		}

		recursivelyGetSequences(s.Values)
	}

	return strconv.Itoa(sum)
}