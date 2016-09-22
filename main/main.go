package main

import (
	"fmt"
	"github.com/abhyudaya/conway/life"
)

func makeGrid(x, y int) [][]uint8 {
	grid := make([][]uint8, x)
	for i := range grid {
		grid[i] = make([]uint8, y)
	}
	return grid
}

func main() {
	x := 10
	y := 10
	n := 25

	current := makeGrid(x, y)
	next := makeGrid(x, y)
	Display(current, "Initial life")

	for i := 1; i <= n; i++ {
		_ = life.NextLife(current, next)
		Display(next, fmt.Sprintf("Step %v", i))
		current, next = next, current
	}

}

func Display(current [][]uint8, msg string) {
	fmt.Println(msg)
	fmt.Println("================")
	for i := range current {
		for j := range current[i] {
			if current[i][j] == 1 {
				fmt.Print(current[i][j])
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
	fmt.Println("================")
}
