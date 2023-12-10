package puzzles

import (
	"advent/utils"
	"strconv"
)


func Puzzle10b() string {
	data := utils.FileReader("data/10.txt")
	pipeMap := buildPipeMap(data)

	start := pipeMap.StartingPositionXY
	from := pipeMap.StartingPositionXY
	compare := pipeMap.GetFirstStep()

	var isLeftOnInside bool

	// Maps routes and records all context
	for {
		// records which directions are inside and outside
		pipeMap.SetContext(from, compare, &isLeftOnInside)
		
		if (start.isSamePosition(compare)) {
			break
		}

		nextPoint := pipeMap.getNextPoint(from, compare)
		from = compare
		compare = nextPoint
	}

	count := 0

	for y := 0; y < len(pipeMap.Grid); y++ {

		checkPoint:
		for x := 0; x < len(pipeMap.Grid[y]); x++ {
			point := &pipeMap.Grid[y][x]

			// ignore pipe parts in loop
			if point.isPipe() && point.isInLoop() {
				continue
			}

			current := point

			// count left from current X until you hit a pipe or the edge
			for pointX := point.X; pointX > 0; pointX-- {
				compare := pipeMap.GetLeftFrom(current)

				if compare.isGround() || (compare.isPipe() && !compare.isInLoop()) {
					current = compare
					continue
				}

				if compare.Context.Right == isLeftOnInside {
					count++
				}
				
				current = compare
				continue checkPoint
			}
		}
	}

	return strconv.Itoa(count)
}