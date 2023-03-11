package main

import (
	"fmt"
	"math"
	"os"

	tetris "tetris/lib"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("ERROR: No input file given")
        fmt.Println("Usage: go run . <input_file>")
        os.Exit(1)
    }
    content := tetris.ReadFileContent(os.Args[1])
    square := Solve(content)
    tetris.PrintGrid(square)
}

func Solve(content string) [][]rune {
	tetrominoes := tetris.GetTetrominoes(content)
	squareSize := int(math.Ceil(math.Sqrt(float64(4 * len(tetrominoes))))) // sqrt(# of tetrominoes * 4 characters per tetromino)
	square := tetris.CreateSquare(squareSize)
	for !tetris.PlaceTetrominoes(square, tetrominoes, 0, 0, 0) {
		squareSize++
		square = tetris.CreateSquare(squareSize)
	}
	return square
}
