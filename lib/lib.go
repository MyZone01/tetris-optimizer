package tetris

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func switchGrid(slice [][]rune) [][]rune {
	result := [][]rune{
		{' ',' ',' ',' ',},
		{' ',' ',' ',' ',},
		{' ',' ',' ',' ',},
		{' ',' ',' ',' ',},
	}
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			result[i][j] = rune(slice[j][i])
		}
	}
	return result
}

func PrintGrid(grid [][]rune) {
	for _, v := range grid {
		for _, w := range v {
			fmt.Print(string(w))
		}
		fmt.Println()
	}
}


func ReadFileContent(file string) string {
	content, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("ERROR: exit when reading file")
		fmt.Println(err)
		os.Exit(1)
	}
	return string(content)
}

func GetTetrominoes(content string) [][][]rune {
	content = strings.ReplaceAll(content, "\r\n", "\n") // TODO: Remove on linux
	_tetrominoes := strings.Split(content, "\n\n")
	tetrominoes := [][][]rune{}
	for i, t := range _tetrominoes {
		_tetromino := strings.Split(t, "\n")
		tetromino := [][]rune{}
		letter := string(rune('A' + i))
		for _, l := range _tetromino {
			if len(l) != 4 {
				fmt.Println("ERROR: Bad format: ",l)
				os.Exit(1)
			}
			tetromino = append(tetromino, []rune(strings.ReplaceAll(l, "#", letter)))
		}
		if len(tetromino) != 4 {
			fmt.Println("ERROR: Bad format: ")
			PrintGrid(tetromino)
			os.Exit(1)
		}
		checkTetromino(tetromino)
		tetromino = moveUpTetromino(tetromino, 0, letter)
		tetromino = moveLeftTetromino(tetromino, 0, letter)
		tetrominoes = append(tetrominoes, tetromino)
	}
	return tetrominoes
}

func moveUpTetromino(lines [][]rune, max int, letter string) [][]rune {
	if !strings.Contains(string(lines[0]), letter) {
		if max == 3 {
			fmt.Println("File invalid")
			os.Exit(1)
		}
		lines[0] = lines[1]
		lines[1] = lines[2]
		lines[2] = lines[3]
		lines[3] = []rune("....")
		moveUpTetromino(lines, max+1, letter)
	}
	return lines
}

func moveLeftTetromino(lines [][]rune, max int, letter string) [][]rune {
	lines = switchGrid(lines)
	lines = moveUpTetromino(lines, max, letter)
	lines = switchGrid(lines)
	return lines
}

func checkTetromino(tetromino [][]rune) {
	numberOfSide, numberOfBlock := 0, 0
	for i, line := range tetromino {
		for j, block := range line {
			if block != '.' {
				numberOfBlock++
				if i-1 >= 0 && tetromino[i-1][j] != '.' { // Up of the block
					numberOfSide++
				}
				if i+1 < 4 && tetromino[i+1][j] != '.' { // Down of the block
					numberOfSide++
				}
				if j-1 >= 0 && tetromino[i][j-1] != '.' { // Left of the block
					numberOfSide++
				}
				if j+1 < 4 && tetromino[i][j+1] != '.' { // Right of the block
					numberOfSide++
				}
			}
		}
	}
	if numberOfBlock != 4 {
		fmt.Println("ERROR: Tetromino must have 4 block.")
		PrintGrid(tetromino)
		os.Exit(1)
	}
	if numberOfSide == 6 || numberOfSide == 8 {
		return
	}
	fmt.Println("ERROR: Tetromino must have 6 sides total touching between the block.")
	PrintGrid(tetromino)
	os.Exit(1)
}

func CreateSquare(squareSize int) [][]rune {
	var grid [][]rune
	for i := 0; i < squareSize; i++ {
		line := []rune{}
		for j := 0; j < squareSize; j++ {
			line = append(line, '.')
		}
		grid = append(grid, line)
	}
	return grid
}

func PlaceTetrominoes(square [][]rune, tetrominoes [][][]rune, row, col, actualTetromino int) bool {
	if actualTetromino == len(tetrominoes) {
		return true
	}

	for i := 0; i < len(square); i++ {
		for j := 0; j < len(square); j++ {
			tetromino := tetrominoes[actualTetromino]
			if Fits(square, tetromino, i, j) {
				Place(square, tetromino, i, j)
				if PlaceTetrominoes(square, tetrominoes, i+1, j+1, actualTetromino+1) {
					return true
				}
				Remove(square, tetromino, i, j)
			}
		}
		col = 0 // reset column index after each row
	}

	return false
}

func Fits(grid [][]rune, tetromino [][]rune, row, col int) bool {
	for i := 0; i < len(tetromino); i++ {
		for j := 0; j < len(tetromino[i]); j++ {
			if tetromino[i][j] != '.' {
				if row+i >= len(grid) || col+j >= len(grid) || grid[row+i][col+j] != '.' {
					return false
				}
			}
		}
	}
	return true
}

func Place(grid [][]rune, tetromino [][]rune, row, col int) {
	for i := 0; i < len(tetromino); i++ {
		for j := 0; j < len(tetromino[i]); j++ {
			if tetromino[i][j] != '.' {
				grid[row+i][col+j] = tetromino[i][j]
			}
		}
	}
}

func Remove(grid [][]rune, tetromino [][]rune, row, col int) {
	for i := 0; i < len(tetromino); i++ {
		for j := 0; j < len(tetromino[i]); j++ {
			if tetromino[i][j] != '.' {
				grid[row+i][col+j] = '.'
			}
		}
	}
}

func Solve(content string) [][]rune {
	tetrominoes := GetTetrominoes(content)
	squareSize := int(math.Ceil(math.Sqrt(float64(4 * len(tetrominoes))))) // sqrt(# of tetrominoes * 4 characters per tetromino)
	square := CreateSquare(squareSize)
	for !PlaceTetrominoes(square, tetrominoes, 0, 0, 0) {
		squareSize++
		square = CreateSquare(squareSize)
	}
	return square
}
