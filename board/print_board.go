package board

import "fmt"

func (b *board) PrintBoard() {
	for _, row := range b.Ships {
		for _, col := range row {
			fmt.Printf("%s ", col)
		}
		fmt.Print("\n")
	}
}
