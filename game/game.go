package game

import (
	"github.com/alvinwilta/go-simple-battleship-sim/board"
	"github.com/alvinwilta/go-simple-battleship-sim/fileparser"
)

type game struct {
	BoardP1    board.Board
	BoardP2    board.Board
	fileparser fileparser.Parser
}

func NewGame(parser fileparser.Parser) Game {
	return &game{
		fileparser: parser,
	}
}

type Game interface {
	// read input file, turn into ListPoint, init board,
	StartGame(inputfile string)
}
