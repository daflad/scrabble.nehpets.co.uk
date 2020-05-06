package controllers

import (
	"github.com/daflad/scrabble.nehpets.co.uk/app/models/pieces"
	"github.com/revel/revel"
)

// App controller
type App struct {
	*revel.Controller
}

// Index page for now
func (c App) Index() revel.Result {
	board := pieces.DefualtBoard()
	var bag pieces.Bag
	bag.Fill()
	rack := bag.RandomRack(7)
	return c.Render(board, bag, rack)
}
