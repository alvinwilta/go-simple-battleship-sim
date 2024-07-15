package main

import (
	"github.com/alvinwilta/gojek-coding-test/fileparser"
	"github.com/alvinwilta/gojek-coding-test/game"
)

func main() {
	inputFilePath := "input.txt"
	fileparser := fileparser.NewParser()
	game := game.NewGame(fileparser)

	game.StartGame(inputFilePath)
}
