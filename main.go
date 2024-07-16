package main

import (
	"github.com/alvinwilta/go-simple-battleship-sim/fileparser"
	"github.com/alvinwilta/go-simple-battleship-sim/game"
)

func main() {
	inputFilePath := "input.txt"
	fileparser := fileparser.NewParser()
	game := game.NewGame(fileparser)

	game.StartGame(inputFilePath)
}
