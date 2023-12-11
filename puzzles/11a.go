package puzzles

import (
	"advent/utils"
	"fmt"
	"math"
	"strconv"
)

type Universe struct {
	Grid [][]string
}

func (u *Universe) expand() {
	expandedY := make([][]string, 0)

	for _, row := range u.Grid {
		hasGalaxy := false

		for _, cell := range row {
			if cell == "#" {
				hasGalaxy = true
			}
		}

		expandedY = append(expandedY, row)

		if hasGalaxy {
			continue
		}

		// no galaxy, add another empty row
		expandedY = append(expandedY, make([]string, len(row)))
	}

	// make new grid and fill
	expanded := make([][]string, len(expandedY))
	for y := range expandedY {
		expanded[y] = make([]string, 0)
	}

	for x := 0; x < len(expandedY[0]); x++ {
		hasGalaxy := false

		for y := 0; y < len(expandedY); y++ {
			if expandedY[y][x] == "#" {
				hasGalaxy = true
			}

			expanded[y] = append(expanded[y], expandedY[y][x])
		}

		if hasGalaxy {
			continue
		}

		// no galaxy, add empty column
		for y := 0; y < len(expandedY); y++ {
			expanded[y] = append(expanded[y], "*")
		}
	}

	// for _, row := range expanded {
	// 	fmt.Println(row)
	// }

	u.Grid = expanded
}

func (u *Universe) getGalaxies() [][]int {
	galaxies := make([][]int, 0)

	for i, row := range u.Grid {
		for j, cell := range row {
			if cell == "#" {
				galaxies = append(galaxies, []int{i, j})
			}
		}
	}

	return galaxies
}

func calculateDistance(g1, g2 []int) int {
	x := math.Abs(float64(g2[0] - g1[0]))
	y := math.Abs(float64(g2[1] - g1[1]))
	return int(x + y)
}

func calculateDistanceSum(galaxies [][]int) int {
	distanceSum := 0

	for i, g1 := range galaxies {
		for j, g2 := range galaxies {
			if i == j {
				continue
			}

			distance := calculateDistance(g1, g2)
			distanceSum += distance
			fmt.Println(g1, g2, distance)
		
		}
	}

	return distanceSum / 2
}

func getUniverse(data []string) Universe {
	galaxyMap := make([][]string, len(data))

	for i, line := range data {
		galaxyMap[i] = make([]string, len(line))

		for j, char := range line {
			galaxyMap[i][j] = string(char)
		}
	}

	return Universe{Grid: galaxyMap}
}

func Puzzle11a() string {
	data := utils.FileReader("data/11.txt")

	universe := getUniverse(data)
	universe.expand()
	// fmt.Println(universe)
	galaxies := universe.getGalaxies()
	sum := calculateDistanceSum(galaxies)

	return strconv.Itoa(sum)
}