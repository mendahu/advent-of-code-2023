package puzzles

import (
	"advent/utils"
	"strconv"
	"strings"
)

func Puzzle01a() string {
	data := utils.FileReader("data/01a.txt")

	val := 0
	
	for _, line := range data {
		chars := strings.Split(line, "")
		var nums []int

		for _, char := range chars {
			num, err := strconv.Atoi(char)
			if (err == nil) {
				nums = append(nums, num)
			}
		}
		calValue := (nums[0] * 10) + nums[len(nums)-1]
		val = val + calValue
	}

	return strconv.Itoa(val)
}