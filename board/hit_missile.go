package board

import "github.com/alvinwilta/go-simple-battleship-sim/types"

func (b *board) HitMissile(input types.ListPoint) {
	for _, point := range input.Points {
		if b.Ships[point.X][point.Y] == "B" {
			b.Ships[point.X][point.Y] = "X"
			b.Scores++
			continue
		}
		b.Ships[point.X][point.Y] = "O"
	}
}
