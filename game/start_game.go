package game

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/alvinwilta/gojek-coding-test/board"
	"github.com/alvinwilta/gojek-coding-test/types"
)

func (g *game) StartGame(inputfile string) {
	// read the file
	input := strings.Split(g.fileparser.Readfile(inputfile), "\n")

	// 0-1 line matrix size
	// create 2 boards with said matrix size
	n, _ := strconv.Atoi(input[0])
	m, _ := strconv.Atoi(input[1])

	// 2-3 line listpoint for ship location
	// init ListPoint for each player, board.FillBoard() for each player board
	g.BoardP1 = board.NewBoard(n, m)
	listPointP1 := types.ListPoint{
		Points: []types.Point{},
	}
	listPointP1.PopulateListPoint(input[2])
	g.BoardP1.FillBoard(listPointP1)

	g.BoardP2 = board.NewBoard(n, m)
	listPointP2 := types.ListPoint{
		Points: []types.Point{},
	}
	listPointP2.PopulateListPoint(input[3])
	g.BoardP2.FillBoard(listPointP2)

	// 5-6 line location of missiles
	// init ListPoint for each missiles, board.HitMissile() for each player board
	listPointMissileP1 := types.ListPoint{
		Points: []types.Point{},
	}
	listPointMissileP1.PopulateListPoint(input[5])
	g.BoardP1.HitMissile(listPointMissileP1)

	listPointMissileP2 := types.ListPoint{
		Points: []types.Point{},
	}
	listPointMissileP2.PopulateListPoint(input[6])
	g.BoardP2.HitMissile(listPointMissileP2)

	// compare the scores for both board by accessing board.Scores
	winMessage := "It is a draw"
	if g.BoardP1.GetScore() > g.BoardP2.GetScore() {
		winMessage = "Player 1 wins"
	}
	if g.BoardP1.GetScore() < g.BoardP2.GetScore() {
		winMessage = "Player 2 wins"
	}

	fmt.Println("Board 1:")
	g.BoardP1.PrintBoard()
	fmt.Println("\nBoard 2:")
	g.BoardP2.PrintBoard()
	fmt.Printf("\nScore P1: %d \n", g.BoardP1.GetScore())
	fmt.Printf("Score P2: %d \n", g.BoardP2.GetScore())
	fmt.Println(winMessage)

}
