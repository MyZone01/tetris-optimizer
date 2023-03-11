package tetris

import (
	"fmt"
	"os"
	"strings"
)

func ReadFileContent(file string) string {
	content, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("Input file not found")
		os.Exit(1)
	}
	return string(content)
}

func GetTetrominoes(content string) [][][]rune {
	_tetrominoes := strings.Split(content, "\r\n\r\n")
	tetrominoes := [][][]rune{}
	for i, t := range _tetrominoes {
		_tetromino := strings.Split(t, "\r\n")
		tetromino := [][]rune{}
		letter := string('A' + i)
		for _, l := range _tetromino {
			tetromino = append(tetromino, []rune(strings.ReplaceAll(l, "#", letter)))
		}
		CheckTetromino(tetromino)
		tetromino = moveUpTetromino(tetromino, 0, letter)
		tetromino = moveLeftTetromino(tetromino, 0, letter)
		tetrominoes = append(tetrominoes, tetromino)
	}
	for _, t := range tetrominoes {
		PrintGrid(t)
	}
	return tetrominoes
}

func moveUpTetromino(lines [][]rune, max int, letter string) [][]rune {
	if !strings.Contains(string(lines[0]), letter) {
		if max == 3 {
			fmt.Println("File invalid")
			os.Exit(1)
		}
		temp := lines[0]
		lines[0] = lines[1]
		lines[1] = lines[2]
		lines[2] = lines[3]
		lines[3] = temp
		moveUpTetromino(lines, max+1, letter)
	}
	return lines
}

func moveLeftTetromino(lines [][]rune, max int, letter string) [][]rune {
	lines = SwitchGrid(lines)
	lines = moveUpTetromino(lines, max, letter)
	lines = SwitchGrid(lines)
	return lines
}

func CheckTetromino(tetromino [][]rune) {
	numberOfSide, numberOfBlock := 0, 0

	for i, line := range tetromino {
		for j, block := range line {
			if block != '.' {
				numberOfBlock++
				if i+1 < 4 && tetromino[i+1][j] != '.' { // Down of the block
					numberOfSide++
				}
				if i-1 >= 0 && tetromino[i-1][j] != '.' { // Up of the block
					numberOfSide++
				}
				if j+1 < 4 && tetromino[i][j+1] != '.' { // Right of the block
					numberOfSide++
				}
				if j-1 >= 0 && tetromino[i][j-1] != '.' { // Left of the block
					numberOfSide++
				}
			}
		}
	}
	if numberOfBlock != 4 {
		fmt.Println("ERROR: Tetromino must have 4 block.")
		os.Exit(1)
	}
	if numberOfSide == 6 || numberOfSide == 8 {
		return
	}
	fmt.Println("ERROR: Tetromino must have 6 sides total touching between the block.")
	os.Exit(1)
}


