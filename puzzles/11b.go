package puzzles

import (
	"advent/utils"
	"fmt"
	"math"
	"strconv"
)

func (u *Universe) getExpansionBoundaries() (rowBoundaries, colBoundaries []int) {
	rowBoundaries = make([]int, 0)
	colBoundaries = make([]int, 0)

	// get row boundaries
	for row := 0; row < len(u.Grid[0]); row++ {
		rowIsBoundary := true

		for col := 0; col < len(u.Grid); col++ {
			if u.Grid[row][col] == "#" {
				rowIsBoundary = false
			}
		}

		if rowIsBoundary {
			rowBoundaries = append(rowBoundaries, row)
		}
	}

	// get col boundaries
	for col := 0; col < len(u.Grid); col++ {
		colIsBoundary := true

		for row := 0; row < len(u.Grid[0]); row++ {
			if u.Grid[row][col] == "#" {
				colIsBoundary = false
			}
		}

		if colIsBoundary {
			colBoundaries = append(colBoundaries, col)
		}
	}

	return
}


func calculateDistanceOverBoundaries(g1, g2 []int, rowBoundaries []int, colBoundaries []int, boundaryVal int) int {
	fmt.Println(g1, g2)
	x := int(math.Abs(float64(g2[0] - g1[0])))

	for _, boundary := range rowBoundaries {
		var greater int
		var lesser int

		if (g2[0] > g1[0]) {
			greater = g2[0]
			lesser = g1[0]
		} else {
			greater = g1[0]
			lesser = g2[0]
		}

		if lesser < boundary && greater > boundary {
			x += boundaryVal - 1
		}
	}

	y := int(math.Abs(float64(g2[1] - g1[1])))

	for _, boundary := range colBoundaries {
		var greater int
		var lesser int

		if (g2[1] > g1[1]) {
			greater = g2[1]
			lesser = g1[1]
		} else {
			greater = g1[1]
			lesser = g2[1]
		}

		if lesser < boundary && greater > boundary {
			y += boundaryVal - 1
		}
	}

	return x + y
}

func calculateDistanceSumOverBoundaries(galaxies [][]int, rowBoundaries []int, colBoundaries []int, boundaryVal int) int {
	distanceSum := 0

	for i := 0; i < len(galaxies); i++ {
		for j := 1 + i; j < len(galaxies); j++ {
			distance := calculateDistanceOverBoundaries(galaxies[i], galaxies[j], rowBoundaries, colBoundaries, boundaryVal)
			distanceSum += distance
		}
	}

	return distanceSum
}

func Puzzle11b() string {
	data := utils.FileReader("data/11.txt")

	universe := getUniverse(data)
	rowBoundaries, colBoundaries := universe.getExpansionBoundaries()
	galaxies := universe.getGalaxies()
	sum := calculateDistanceSumOverBoundaries(galaxies, rowBoundaries, colBoundaries, 1000000)

	return strconv.Itoa(sum)
}