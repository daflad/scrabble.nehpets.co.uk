package pieces

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/revel/revel"
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

var fullBag = [][]interface{}{
	{"E", 12, 1}, {"A", 9, 1}, {"I", 9, 1}, {"O", 8, 1}, {"N", 6, 1}, {"R", 6, 1},
	{"T", 6, 1}, {"L", 4, 1}, {"S", 4, 1}, {"U", 4, 1}, {"D", 4, 2}, {"G", 3, 2},
	{"B", 2, 3}, {"C", 2, 3}, {"M", 2, 3}, {"M", 2, 3}, {"P", 2, 3}, {"F", 2, 4},
	{"H", 2, 4}, {"V", 2, 4}, {"W", 2, 4}, {"Y", 2, 4}, {"K", 1, 5}, {"J", 1, 8},
	{"X", 1, 8}, {"Q", 1, 10}, {"Z", 1, 10},
}

// Letter for playing the game
type Letter struct {
	Name     string
	Value    int
	Position Location
}

func (l Letter) String() string {
	return fmt.Sprintf("%s", l.Name)
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

func letterFromSet(set []interface{}) (l Letter) {
	l.Name = set[0].(string)
	l.Value = set[2].(int)
	return
}

// Fill the bag with default letter sets
func (b *Bag) Fill() {
	b.Letters = make([]Letter, 0)
	for _, set := range fullBag {
		for i := 0; i < set[1].(int); i++ {
			b.Letters = append(b.Letters, letterFromSet(set))
		}
	}
}

// RandomRack build a rack of random letters from the bag
func (b *Bag) RandomRack(length int) (rack []Letter) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < length; i++ {
		index := rand.Intn(len(b.Letters))
		rack = append(rack, b.Letters[index])
		b.Letters = append(b.Letters[:index], b.Letters[index+1:]...)
	}
	revel.AppLog.Infof("%+v", rack)
	return rack
}
