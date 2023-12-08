package puzzles

import (
	"advent/utils"
	"strconv"
)

// get greatest common divisor
func GCD(a int, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// get least common multiple
func LCM(a int, b int) int {
	result := a * b / GCD(a, b)

	return result
}

func Puzzle08b() string {
	data := utils.FileReader("data/08.txt")
	directions, waypoints, startingNodes := getDesertWaypointData(data)

	var navigate func()

	for node := range startingNodes {
		currentWaypoint := node

		navigate = func () {
			for _, direction := range directions {
				startingNodes[node]++
				currentWaypoint = waypoints[currentWaypoint][direction]

				if (string([]rune(currentWaypoint)[2]) == "Z") {
					return
				}
			}
	
			navigate()
		}
	
		navigate()
	}	

	count := 0

	for _, multiple := range startingNodes {
		if (count == 0) {
			count = multiple
			continue
		}
		count = LCM(count, multiple)
	}

	return strconv.Itoa(count)
}