package game

import (
	"github.com/alvinwilta/gojek-coding-test/board"
	"github.com/alvinwilta/gojek-coding-test/fileparser"
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
