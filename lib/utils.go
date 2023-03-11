package tetris

import "fmt"

func ForEach(f func(rune), a [][]rune) {
	for _, v := range a {
		for _, b := range v {
			f(b)
		}
	}
}

func SwitchGrid(slice [][]rune) [][]rune {
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

func PrintGrid(n [][]rune) {
	for _, v := range n {
		for _, w := range v {
			fmt.Print(string(w))
		}
		fmt.Println()
	}
}
