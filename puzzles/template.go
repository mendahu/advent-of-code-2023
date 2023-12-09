package puzzles

import (
	"advent/utils"
	"fmt"
	"strconv"
)

func Puzzle0Xa() string {
	data := utils.FileReader("data/0X_test.txt")

	fmt.Println(data[0])

	return strconv.Itoa(0)
}