package puzzles

import (
	"advent/utils"
	"fmt"
	"strconv"
)

type PipeContext struct {
	Above bool
	Below bool
	Left bool
	Right bool
	PartOfLoop bool
}

type PipeMapPoint struct {
	Char string
	X int
	Y int
	Context PipeContext
}

func (p *PipeMapPoint) isSamePosition(point *PipeMapPoint) bool {
	if (point == nil || p == nil) {
		return false
	}
	return p.X == point.X && p.Y == point.Y
}

func (p *PipeMapPoint) isGround() bool {
	return p.Char == "."
}

func (p *PipeMapPoint) isPipe() bool {
	return !p.isGround()
}

func (p *PipeMapPoint) isInLoop() bool {
	return p.Context.PartOfLoop
}

type PipeMap struct {
	Grid [][]PipeMapPoint
	StartingPositionXY *PipeMapPoint
}


func (p *PipeMap) GetFirstStep() *PipeMapPoint {
	start := p.StartingPositionXY

	if start.Char == "F" {
		return p.GetRightFrom(start)
	}

	if start.Char == "7" {
		return p.GetLeftFrom(start)
	}

	if start.Char == "|" {
		return p.GetAboveFrom(start)
	}

	if start.Char == "-" {
		return p.GetRightFrom(start)
	}

	if start.Char == "J" {
		return p.GetAboveFrom(start)
	}

	if start.Char == "L" {
		return p.GetBelowFrom(start)
	}

	return nil
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
	if (from.X == 0) {
		return nil
	}
	return &p.Grid[from.Y][from.X - 1]
}

func (p *PipeMap) GetRightFrom(from *PipeMapPoint) *PipeMapPoint {
	if (from.X == len(p.Grid[0]) - 1) {
		return nil
	}
	return &p.Grid[from.Y][from.X + 1]
}

func (p *PipeMap) GetAboveFrom(from *PipeMapPoint) *PipeMapPoint {
	if (from.Y == 0) {
		return nil
	}
	return &p.Grid[from.Y - 1][from.X]
}

func (p *PipeMap) GetBelowFrom(from *PipeMapPoint) *PipeMapPoint {
	if (from.Y == len(p.Grid) - 1) {
		return nil
	}
	return &p.Grid[from.Y + 1][from.X]
}

func (p *PipeMap) SetContext(from *PipeMapPoint, current *PipeMapPoint, isLeftOnInside *bool) {
	current.Context.PartOfLoop = true
	
	// We arbitrarily set Left of the pipe track as true. Later we will asign true or false as inside or outside

	// came from left
	if (p.GetLeftFrom(current).isSamePosition(from)) {
		if current.Char == "7" {
			current.Context.Above = true
			current.Context.Right = true
		}

		if current.Char == "-" {
			current.Context.Above = true
		}

		if (current.Y == len(p.Grid) - 1) {
			*isLeftOnInside = true
		}
		return;
	}

	//came from above
	if (p.GetAboveFrom(current).isSamePosition(from)) {
		if current.Char == "J" {
			current.Context.Right = true
			current.Context.Below = true
		}

		if current.Char == "|" {
			current.Context.Right = true
		}

		if (current.X == 0) {
			*isLeftOnInside = true
		}
		return;
	}

	// came from right
	if (p.GetRightFrom(current).isSamePosition(from)) {
		if current.Char == "-" {
			current.Context.Below = true
		}

		if current.Char == "L" {
			current.Context.Below = true
			current.Context.Left = true
		}

		if (current.Y == 0) {
			*isLeftOnInside = true
		}
		return;
	}

	// came from below
	if (p.GetBelowFrom(current).isSamePosition(from)) {
		if current.Char == "|" {
			current.Context.Left = true
		}

		if current.Char == "F" {
			current.Context.Left = true
			current.Context.Above = true
		}

		if (current.X == len(p.Grid[0]) - 1) {
			*isLeftOnInside = true
		}
		return;
	}
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
				Context: PipeContext{},
			}

			if (char == "S") {
				startingPositionXY = &point
				point.Context.PartOfLoop = true
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
	compare := pipeMap.GetFirstStep()

	for {
		steps++
		
		if (start.isSamePosition(compare)) {
			break
		}

		nextPoint := pipeMap.getNextPoint(from, compare)
		from = compare
		compare = nextPoint
	}

	midpoint := steps / 2

	return strconv.Itoa(midpoint)
}