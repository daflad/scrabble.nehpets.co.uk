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
	for i := 0; i < 4; i++ {
		rack := bag.RandomRack(7)
		for j, letter := range rack {
			board.Tiles[i][j].Value = letter.Name
			board.Tiles[i][j].Points = letter.Value
		}
	}
	i, j := 5, 0
	for _, letter := range bag.Letters {
		board.Tiles[i][j].Value = letter.Name
		board.Tiles[i][j].Points = letter.Value
		j++
		if j > 14 {
			j = 0
			i++
		}
	}
	return c.Render(board, bag)
}
