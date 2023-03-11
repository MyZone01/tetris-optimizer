package main

import (
	"fmt"
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
    square := tetris.Solve(content)
    tetris.PrintGrid(square)
}
