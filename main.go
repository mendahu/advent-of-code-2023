package main

import (
	"advent/lib"
	"fmt"
	"os"
)

// go run main.go 01a

func main() {
	code := os.Args[1]

	puzzle := lib.GetPuzzleFunc(code)
	answer := puzzle()
	fmt.Println(answer)
}