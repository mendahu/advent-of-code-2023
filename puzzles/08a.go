package puzzles

import (
	"advent/utils"
	"fmt"
	"strconv"
	"strings"
)

func getDesertWaypointData(data []string) (directions []int, waypoints map[string][]string, startingNodes map[string]int) {
	directions = []int{}
	waypoints = map[string][]string{}
	startingNodes = map[string]int{}
	
	directionsString := data[0]

	for _, direction := range directionsString {
		if direction == 'L' {
			directions = append(directions, 0)
		} else {
			directions = append(directions, 1)
		}
	}

	for _, waypoint := range data[2:] {
		var start string
		var left string
		var right string

		r := strings.NewReader(waypoint)
		fmt.Fscanf(r, "%s = (%3s, %3s)", &start, &left, &right)
		waypoints[start] = []string{left, right}

		if string([]rune(start)[2]) == "A" {
			startingNodes[start] = 0
		}
	}

	return
}

func Puzzle08a() string {
	data := utils.FileReader("data/08.txt")
	directions, waypoints, _ := getDesertWaypointData(data)
	fmt.Println(directions)
	fmt.Println(waypoints)
	count := 0

	currentWaypoint := "AAA"

	var navigate func()

	navigate = func () {
		for _, direction := range directions {
			count++
			
			currentWaypoint = waypoints[currentWaypoint][direction]

			if (currentWaypoint == "ZZZ") {
				return
			}
		}

		navigate()
	}

	navigate()

	return strconv.Itoa(count)
}