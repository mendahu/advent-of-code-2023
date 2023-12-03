package puzzles

import (
	"advent/utils"
	"strconv"
)

func isNumber(char string) bool {
	if char == "1" || char == "2" || char == "3" || char == "4" || char == "5" || char == "6" || char == "7" || char == "8" || char == "9" || char == "0" {
		return true
	}

	return false
}

func isIndicator(char string) bool {
	if (isNumber(char)) {
		return false
	}
	if char == "." {
		return false
	}

	return true
}

type Indicator struct {
	Index int
	Gear bool
	PartNumbers []int
}

type PartNumber struct {
	PartNumber int
	Indices []int
	Counted bool
}

type Line struct {
	Indicators  []Indicator
	PartNumbers []PartNumber
	Data        string
}

func (l *Line) getSymbolIndices() {
	l.Indicators = []Indicator{}

	for i := 0; i < len(l.Data); i++ {
		char := string(l.Data[i])
		if isIndicator(char) {
			isGear := char == "*"
			l.Indicators = append(l.Indicators, Indicator{Index: i, Gear: isGear, PartNumbers: make([]int, 0)})
		}
	}
}

func (l *Line) getPartNumbers() {
	l.PartNumbers = []PartNumber{}

	i := 0
	for {
		if i >= len(l.Data) {
			break
		}

		if (!isNumber(string(l.Data[i]))) {
			i++
			continue
		}

		num := string(l.Data[i])
		length := 1
		partNumber := PartNumber{
			Indices: []int{i},
			Counted: false,
		}

		for {
			if (i + length == len(l.Data) || !isNumber(string(l.Data[i + length]))) {
				numInt, err := strconv.Atoi(num)
				if err != nil {
					panic(err)
				}
				partNumber.PartNumber = numInt
				l.PartNumbers = append(l.PartNumbers, partNumber)
				i += length
				break
			}

			num += string(l.Data[i + length])
			partNumber.Indices = append(partNumber.Indices, i + length)
			length++
		}
	}
}

func (l *Line) initialize() {
	l.getSymbolIndices()
	l.getPartNumbers()
}

func Puzzle03a() string {
	data := utils.FileReader("data/03.txt")

	var preceding Line
	var current Line
	var following Line

	lengOfData := len(data[0])
	baseString := ""
	for i := 0; i < lengOfData; i++ {
		baseString += "."
	}

	preceding = Line{Data: baseString}
	current = Line{Data: data[0]}

	preceding.initialize()
	current.initialize()

	sum := 0

	for i := 0; i < len(data); i++ {
		if (i + 1 == len(data)) {
			following = Line{Data: baseString}
		} else {
			following = Line{Data: data[i+1]}
		}

		following.initialize()

		for _, indicator := range current.Indicators {
			for _, line := range []Line{preceding, current, following} {

				checkpart:
				for _, partNumber := range line.PartNumbers {
					if partNumber.Counted {
						continue
					}

					for _, partNumberIndex := range partNumber.Indices {
						if partNumberIndex == indicator.Index - 1 || partNumberIndex == indicator.Index || partNumberIndex == indicator.Index + 1  {
							sum += partNumber.PartNumber
							partNumber.Counted = true
							continue checkpart
						}
					}
				}
			}
		}
		
		preceding = current
		current = following
	}

	

	return strconv.Itoa(sum)
}