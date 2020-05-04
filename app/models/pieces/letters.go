package pieces

import (
	"math/rand"
	"time"
)

// Letter distribution & points
//
// 1 point: E ×12, A ×9, I ×9, O ×8, N ×6, R ×6, T ×6, L ×4, S ×4, U ×4
// 2 points: D ×4, G ×3
// 3 points: B ×2, C ×2, M ×2, P ×2
// 4 points: F ×2, H ×2, V ×2, W ×2, Y ×2
// 5 points: K ×1
// 8 points: J ×1, X ×1
// 10 points: Q ×1, Z ×1

// Letter for playing the game
type Letter struct {
	Name     string
	Value    int
	Position Location
}

// Location of a letter on the board
type Location struct {
	X int
	Y int
}

// Bag of letters to pick from
type Bag struct {
	Letters []Letter
}

// RandomRack build a rack of random letters from the bag
func (b Bag) RandomRack(length int) (rack []Letter) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < length; i++ {
		index := rand.Intn(len(b.Letters))
		rack = append(rack, b.Letters[index])
		b.Letters = append(b.Letters[:index], b.Letters[index+1:]...)
	}
	return rack
}
