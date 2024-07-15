package board

import "github.com/alvinwilta/gojek-coding-test/types"

type board struct {
	// matrix
	Ships  map[int]map[int]string
	Scores int
}

func NewBoard(n int, m int) Board {
	boardShip := make(map[int]map[int]string, n)
	for i := 0; i < n; i++ {
		boardShip[i] = make(map[int]string, m)
		for j := 0; j < m; j++ {
			boardShip[i][j] = "_"
		}
	}
	return &board{
		Ships:  boardShip,
		Scores: 0,
	}
}

type Board interface {
	// fill the boar dwith appropriate ship location
	FillBoard(input types.ListPoint)

	// calculate the missiles on the board
	HitMissile(input types.ListPoint)

	// get scores
	GetScore() int
}
