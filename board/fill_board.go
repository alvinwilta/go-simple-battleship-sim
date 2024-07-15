package board

import "github.com/alvinwilta/gojek-coding-test/types"

func (b *board) FillBoard(input types.ListPoint) {
	for _, point := range input.Points {
		b.Ships[point.X][point.Y] = "B"
	}
}
