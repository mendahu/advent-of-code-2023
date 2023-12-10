package puzzles

import (
	"advent/utils"
	"fmt"
	"strconv"
)

type PipeMapPoint struct {
	Char string
	X int
	Y int
}

func (p *PipeMapPoint) isSamePosition(point *PipeMapPoint) bool {
	fmt.Println(p.X, p.Y, point.X, point.Y)
	return p.X == point.X && p.Y == point.Y
}

type PipeMap struct {
	Grid [][]PipeMapPoint
	StartingPositionXY *PipeMapPoint
}

func (p *PipeMap) SetStartPositionChar() {
	var char string
	startingPoint := p.StartingPositionXY
	height := len(p.Grid)
	width := len(p.Grid[0])

	hasAbove := false
	hasBelow := false
	hasLeft := false
	hasRight := false
	
	// Get first direction to attempts, starting from top and rotating around clockwise
	// Check Above
	if (startingPoint.Y > 0) {
		above := p.Grid[startingPoint.Y - 1][startingPoint.X]
		if (above.Char == "|" || above.Char == "F" || above.Char == "7") {
			hasAbove = true
		}
	}

	// Check Right
	if (startingPoint.X < width) {
		right := p.Grid[startingPoint.Y][startingPoint.X + 1]
		if (right.Char == "-" || right.Char == "J" || right.Char == "7") {
			hasRight = true
		}
	}

	// Check Below
	if (startingPoint.Y < height) {
		below := p.Grid[startingPoint.Y + 1][startingPoint.X]
		if (below.Char == "|" || below.Char == "L" || below.Char == "J") {
			hasBelow = true
		}
	}

	// Check Left
	if (startingPoint.X > 0) {
		left := p.Grid[startingPoint.Y][startingPoint.X - 1]
		if (left.Char == "-" || left.Char == "F" || left.Char == "L") {
			hasLeft = true
		}
	}

	if (hasAbove && hasLeft) {
		char = "J"
	} else if (hasAbove && hasRight) {
		char = "L"
	} else if (hasBelow && hasLeft) {
		char = "7"
	} else if (hasBelow && hasRight) {
		char = "F"
	} else if (hasAbove && hasBelow) {
		char = "|"
	} else if (hasLeft && hasRight) {
		char = "-"
	} else {
		fmt.Println("ERROR: Could not determine starting position char")
	}

	startingPoint.Char = char
}

func (p *PipeMap) GetLeftFrom(from *PipeMapPoint) *PipeMapPoint {
	return &p.Grid[from.Y][from.X - 1]
}

func (p *PipeMap) GetRightFrom(from *PipeMapPoint) *PipeMapPoint {
	return &p.Grid[from.Y][from.X + 1]
}

func (p *PipeMap) GetAboveFrom(from *PipeMapPoint) *PipeMapPoint {
	return &p.Grid[from.Y - 1][from.X]
}

func (p *PipeMap) GetBelowFrom(from *PipeMapPoint) *PipeMapPoint {
	return &p.Grid[from.Y + 1][from.X]
}

func (p *PipeMap) getNextPoint(from *PipeMapPoint, current *PipeMapPoint) *PipeMapPoint {
	// if from is below, then check the char for left, right or above
	if (from.Y > current.Y && from.X == current.X) {
		// Check left
		if (current.Char == "7") {
			return p.GetLeftFrom(current)
		}

		// Check right
		if (current.Char == "F") {
			return p.GetRightFrom(current)
		}

		// Check above
		if (current.Char == "|") {
			return p.GetAboveFrom(current)
		}
	}

	// if from is above, then check the char for left, right or below
	if (from.Y < current.Y && from.X == current.X) {
		// Check left
		if (current.Char == "J") {
			return p.GetLeftFrom(current)
		}

		// Check right
		if (current.Char == "L") {
			return p.GetRightFrom(current)
		}

		// Check below
		if (current.Char == "|") {
			return p.GetBelowFrom(current)
		}
	}

	// if from is left, then check the char for above, below or right
	if (from.X < current.X && from.Y == current.Y) {
		// Check above
		if (current.Char == "J") {
			return p.GetAboveFrom(current)
		}

		// Check below
		if (current.Char == "7") {
			return p.GetBelowFrom(current)
		}

		// Check right
		if (current.Char == "-") {
			return p.GetRightFrom(current)
		}
	}

	// if from is right, then check the char for above, below or left
	if (from.X > current.X && from.Y == current.Y) {
		// Check above
		if (current.Char == "L") {
			return p.GetAboveFrom(current)
		}

		// Check below
		if (current.Char == "F") {
			return p.GetBelowFrom(current)
		}

		// Check left
		if (current.Char == "-") {
			return p.GetLeftFrom(current)
		}
	}

	return nil
}

func buildPipeMap(rows []string) PipeMap {
	grid := make([][]PipeMapPoint, len(rows))
	var startingPositionXY *PipeMapPoint

	for y, line := range rows {
		grid[y] = make([]PipeMapPoint, len(line))

		for x, char := range line {
			char := string(char)
			point := PipeMapPoint{
				Char: char,
				X: x,
				Y: y,
			}

			if (char == "S") {
				startingPositionXY = &point
			}

			grid[y][x] = point
		}
	}

	pipeMap := PipeMap{Grid: grid, StartingPositionXY: startingPositionXY}
	pipeMap.SetStartPositionChar()

	return pipeMap
}

func Puzzle10a() string {
	data := utils.FileReader("data/10.txt")
	pipeMap := buildPipeMap(data)

	steps := 0
	start := pipeMap.StartingPositionXY
	from := pipeMap.StartingPositionXY
	var compare *PipeMapPoint

	if (start.Char == "F" || start.Char == "L") {
		compare = pipeMap.GetRightFrom(start)
	} else if (start.Char == "-" || start.Char == "J") {
		compare = pipeMap.GetLeftFrom(start)
	} else {
		compare = pipeMap.GetAboveFrom(start)
	}

	for {
		steps++
		
		if (start.isSamePosition(compare)) {
			break
		}

		fmt.Println(from, compare)
		nextPoint := pipeMap.getNextPoint(from, compare)
		from = compare
		compare = nextPoint
	}

	midpoint := steps / 2

	return strconv.Itoa(midpoint)
}