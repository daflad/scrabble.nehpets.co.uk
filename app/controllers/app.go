package controllers

import (
	"io/ioutil"
	"strings"

	"github.com/daflad/scrabble.nehpets.co.uk/app/models/pieces"
	"github.com/revel/revel"
)

// App controller
type App struct {
	*revel.Controller
}

// Game page for now
func (c App) Game() revel.Result {
	board := pieces.DefualtBoard()
	var bag pieces.Bag
	bag.Fill()
	rack := bag.RandomRack(7)
	bag.RandomRack(7)
	bag.RandomRack(7)
	bag.RandomRack(7)
	bag.RandomRack(5)
	return c.Render(board, bag, rack)
}

var scores = map[string]int{
	"E": 1, "A": 1, "I": 1, "O": 1, "N": 1, "R": 1, "T": 1, "L": 1, "S": 1,
	"U": 1, "D": 2, "G": 2, "B": 3, "C": 3, "M": 3, "P": 3, "F": 4, "H": 4,
	"V": 4, "W": 4, "Y": 4, "K": 5, "J": 8, "X": 8, "Q": 10, "Z": 10,
}

// Index page for now
func (c App) Index() revel.Result {
	board := pieces.DefualtBoard()
	dat, err := ioutil.ReadFile(revel.AppRoot + "scrabble.csv")
	if err != nil {
		revel.AppLog.Error(err.Error())
	}
	for i, line := range strings.Split(string(dat), "\n") {
		for j, cell := range strings.Split(line, ";") {
			cell = strings.Replace(cell, "\"", "", -1)
			board.Tiles[i][j].Value = cell
			board.Tiles[i][j].Points = scores[cell]
		}
	}
	return c.Render(board)
}
