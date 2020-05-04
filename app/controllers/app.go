package controllers

import (
	"github.com/daflad/scrabble.nehpets.co.uk/app/models/pieces"
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	board := pieces.DefualtBoard()
	board.Tiles[1][6].Value = "J"
	board.Tiles[1][6].Points = 8
	board.Tiles[1][7].Value = "O"
	board.Tiles[1][7].Points = 1
	return c.Render(board)
}
