package board

import "github.com/alvinwilta/go-simple-battleship-sim/types"

func (b *board) FillBoard(input types.ListPoint) {
	for _, point := range input.Points {
		b.Ships[point.X][point.Y] = "B"
	}
}
