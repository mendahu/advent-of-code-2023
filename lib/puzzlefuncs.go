package lib

import "advent/puzzles"

func GetPuzzleFunc(code string) func()string {
	puzzles := map[string]func()string{
		"01a": puzzles.Puzzle01a,
		"01b": puzzles.Puzzle01b,
		"02a": puzzles.Puzzle02a,
		"02b": puzzles.Puzzle02b,
	}

	return puzzles[code]
}