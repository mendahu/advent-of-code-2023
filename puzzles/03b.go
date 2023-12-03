package puzzles

import (
	"advent/utils"
	"strconv"
)

func Puzzle03b() string {
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

			if !indicator.Gear {
				continue
			}

			for _, line := range []Line{preceding, current, following} {

				checkpart:
				for _, partNumber := range line.PartNumbers {
					for _, partNumberIndex := range partNumber.Indices {
						if partNumberIndex == indicator.Index - 1 || partNumberIndex == indicator.Index || partNumberIndex == indicator.Index + 1  {
							indicator.PartNumbers = append(indicator.PartNumbers, partNumber.PartNumber)
							continue checkpart
						}
					}
				}
			}

			if len(indicator.PartNumbers) == 2 {
				sum += indicator.PartNumbers[0] * indicator.PartNumbers[1]
			}
		}
		
		preceding = current
		current = following
	}

	return strconv.Itoa(sum)
}
