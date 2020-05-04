package pieces

// Board the scrabble board to play on
type Board struct {
	Tiles [15][15]Tile
}

// Tile the individual tiles of a scrabble board
type Tile struct {
	Type   TileType
	CSS    string
	Value  string
	Points int
}

// TileType enum
type TileType int

// Tiletypes enums
const (
	Tiletypes = iota
	DefualtTile
	DoubleLetter
	TrippleLetter
	DoubleWord
	TrippleWord
)

// newTile for DefualtBoard
func newTile(tt TileType) (t Tile) {
	t.Type = tt
	t.CSS = cssClass(tt)
	return
}

func cssClass(tt TileType) string {
	switch tt {
	case 2:
		return "double-letter"
	case 3:
		return "tripple-letter"
	case 4:
		return "double-word"
	case 5:
		return "tripple-word"
	default:
		return "default-tile"
	}
}

func (t TileType) String() string {
	names := []string{
		"DEFAULT TILE",
		"DOUBLE LETTER",
		"TRIPPLE LETTER",
		"DOUBLE WORD",
		"TRIPPLE WORD",
	}
	return names[t-1]
}

// DefualtBoard construct the empty board
func DefualtBoard() (b Board) {
	// Row 1
	b.Tiles[0][0] = newTile(TrippleWord)
	b.Tiles[0][1] = newTile(DefualtTile)
	b.Tiles[0][2] = newTile(DefualtTile)
	b.Tiles[0][3] = newTile(DoubleLetter)
	b.Tiles[0][4] = newTile(DefualtTile)
	b.Tiles[0][5] = newTile(DefualtTile)
	b.Tiles[0][6] = newTile(DefualtTile)
	b.Tiles[0][7] = newTile(TrippleWord)
	b.Tiles[0][8] = newTile(DefualtTile)
	b.Tiles[0][9] = newTile(DefualtTile)
	b.Tiles[0][10] = newTile(DefualtTile)
	b.Tiles[0][11] = newTile(DoubleLetter)
	b.Tiles[0][12] = newTile(DefualtTile)
	b.Tiles[0][13] = newTile(DefualtTile)
	b.Tiles[0][14] = newTile(TrippleWord)

	// Row 2
	b.Tiles[1][0] = newTile(DefualtTile)
	b.Tiles[1][1] = newTile(DoubleWord)
	b.Tiles[1][2] = newTile(DefualtTile)
	b.Tiles[1][3] = newTile(DefualtTile)
	b.Tiles[1][4] = newTile(DefualtTile)
	b.Tiles[1][5] = newTile(TrippleLetter)
	b.Tiles[1][6] = newTile(DefualtTile)
	b.Tiles[1][7] = newTile(DefualtTile)
	b.Tiles[1][8] = newTile(DefualtTile)
	b.Tiles[1][9] = newTile(TrippleLetter)
	b.Tiles[1][10] = newTile(DefualtTile)
	b.Tiles[1][11] = newTile(DefualtTile)
	b.Tiles[1][12] = newTile(DefualtTile)
	b.Tiles[1][13] = newTile(DoubleWord)
	b.Tiles[1][14] = newTile(DefualtTile)

	// Row 3
	b.Tiles[2][0] = newTile(DefualtTile)
	b.Tiles[2][1] = newTile(DefualtTile)
	b.Tiles[2][2] = newTile(DoubleWord)
	b.Tiles[2][3] = newTile(DefualtTile)
	b.Tiles[2][4] = newTile(DefualtTile)
	b.Tiles[2][5] = newTile(DefualtTile)
	b.Tiles[2][6] = newTile(DoubleLetter)
	b.Tiles[2][7] = newTile(DefualtTile)
	b.Tiles[2][8] = newTile(DoubleLetter)
	b.Tiles[2][9] = newTile(DefualtTile)
	b.Tiles[2][10] = newTile(DefualtTile)
	b.Tiles[2][11] = newTile(DefualtTile)
	b.Tiles[2][12] = newTile(DoubleWord)
	b.Tiles[2][13] = newTile(DefualtTile)
	b.Tiles[2][14] = newTile(DefualtTile)

	// Row 4
	b.Tiles[3][0] = newTile(DoubleLetter)
	b.Tiles[3][1] = newTile(DefualtTile)
	b.Tiles[3][2] = newTile(DefualtTile)
	b.Tiles[3][3] = newTile(DoubleWord)
	b.Tiles[3][4] = newTile(DefualtTile)
	b.Tiles[3][5] = newTile(DefualtTile)
	b.Tiles[3][6] = newTile(DefualtTile)
	b.Tiles[3][7] = newTile(DoubleLetter)
	b.Tiles[3][8] = newTile(DefualtTile)
	b.Tiles[3][9] = newTile(DefualtTile)
	b.Tiles[3][10] = newTile(DefualtTile)
	b.Tiles[3][11] = newTile(DoubleWord)
	b.Tiles[3][12] = newTile(DefualtTile)
	b.Tiles[3][13] = newTile(DefualtTile)
	b.Tiles[3][14] = newTile(DoubleLetter)

	// Row 5
	b.Tiles[4][0] = newTile(DefualtTile)
	b.Tiles[4][1] = newTile(DefualtTile)
	b.Tiles[4][2] = newTile(DefualtTile)
	b.Tiles[4][3] = newTile(DefualtTile)
	b.Tiles[4][4] = newTile(DoubleWord)
	b.Tiles[4][5] = newTile(DefualtTile)
	b.Tiles[4][6] = newTile(DefualtTile)
	b.Tiles[4][7] = newTile(DefualtTile)
	b.Tiles[4][8] = newTile(DefualtTile)
	b.Tiles[4][9] = newTile(DefualtTile)
	b.Tiles[4][10] = newTile(DoubleWord)
	b.Tiles[4][11] = newTile(DefualtTile)
	b.Tiles[4][12] = newTile(DefualtTile)
	b.Tiles[4][13] = newTile(DefualtTile)
	b.Tiles[4][14] = newTile(DefualtTile)

	// Row 6
	b.Tiles[5][0] = newTile(DefualtTile)
	b.Tiles[5][1] = newTile(TrippleLetter)
	b.Tiles[5][2] = newTile(DefualtTile)
	b.Tiles[5][3] = newTile(DefualtTile)
	b.Tiles[5][4] = newTile(DefualtTile)
	b.Tiles[5][5] = newTile(TrippleLetter)
	b.Tiles[5][6] = newTile(DefualtTile)
	b.Tiles[5][7] = newTile(DefualtTile)
	b.Tiles[5][8] = newTile(DefualtTile)
	b.Tiles[5][9] = newTile(TrippleLetter)
	b.Tiles[5][10] = newTile(DefualtTile)
	b.Tiles[5][11] = newTile(DefualtTile)
	b.Tiles[5][12] = newTile(DefualtTile)
	b.Tiles[5][13] = newTile(TrippleLetter)
	b.Tiles[5][14] = newTile(DefualtTile)

	// Row 7
	b.Tiles[6][0] = newTile(DefualtTile)
	b.Tiles[6][1] = newTile(DefualtTile)
	b.Tiles[6][2] = newTile(DoubleLetter)
	b.Tiles[6][3] = newTile(DefualtTile)
	b.Tiles[6][4] = newTile(DefualtTile)
	b.Tiles[6][5] = newTile(DefualtTile)
	b.Tiles[6][6] = newTile(DoubleLetter)
	b.Tiles[6][7] = newTile(DefualtTile)
	b.Tiles[6][8] = newTile(DoubleLetter)
	b.Tiles[6][9] = newTile(DefualtTile)
	b.Tiles[6][10] = newTile(DefualtTile)
	b.Tiles[6][11] = newTile(DefualtTile)
	b.Tiles[6][12] = newTile(DoubleLetter)
	b.Tiles[6][13] = newTile(DefualtTile)
	b.Tiles[6][14] = newTile(DefualtTile)

	// Row 8
	b.Tiles[7][0] = newTile(TrippleWord)
	b.Tiles[7][1] = newTile(DefualtTile)
	b.Tiles[7][2] = newTile(DefualtTile)
	b.Tiles[7][3] = newTile(DoubleLetter)
	b.Tiles[7][4] = newTile(DefualtTile)
	b.Tiles[7][5] = newTile(DefualtTile)
	b.Tiles[7][6] = newTile(DefualtTile)
	b.Tiles[7][7] = newTile(DoubleWord)
	b.Tiles[7][7].CSS += " centre-star"
	b.Tiles[7][8] = newTile(DefualtTile)
	b.Tiles[7][9] = newTile(DefualtTile)
	b.Tiles[7][10] = newTile(DefualtTile)
	b.Tiles[7][11] = newTile(DoubleLetter)
	b.Tiles[7][12] = newTile(DefualtTile)
	b.Tiles[7][13] = newTile(DefualtTile)
	b.Tiles[7][14] = newTile(TrippleLetter)

	// Row 9
	b.Tiles[8][0] = newTile(DefualtTile)
	b.Tiles[8][1] = newTile(DefualtTile)
	b.Tiles[8][2] = newTile(DoubleLetter)
	b.Tiles[8][3] = newTile(DefualtTile)
	b.Tiles[8][4] = newTile(DefualtTile)
	b.Tiles[8][5] = newTile(DefualtTile)
	b.Tiles[8][6] = newTile(DoubleLetter)
	b.Tiles[8][7] = newTile(DefualtTile)
	b.Tiles[8][8] = newTile(DoubleLetter)
	b.Tiles[8][9] = newTile(DefualtTile)
	b.Tiles[8][10] = newTile(DefualtTile)
	b.Tiles[8][11] = newTile(DefualtTile)
	b.Tiles[8][12] = newTile(DoubleLetter)
	b.Tiles[8][13] = newTile(DefualtTile)
	b.Tiles[8][14] = newTile(DefualtTile)

	// Row 10
	b.Tiles[9][0] = newTile(DefualtTile)
	b.Tiles[9][1] = newTile(TrippleLetter)
	b.Tiles[9][2] = newTile(DefualtTile)
	b.Tiles[9][3] = newTile(DefualtTile)
	b.Tiles[9][4] = newTile(DefualtTile)
	b.Tiles[9][5] = newTile(TrippleLetter)
	b.Tiles[9][6] = newTile(DefualtTile)
	b.Tiles[9][7] = newTile(DefualtTile)
	b.Tiles[9][8] = newTile(DefualtTile)
	b.Tiles[9][9] = newTile(TrippleLetter)
	b.Tiles[9][10] = newTile(DefualtTile)
	b.Tiles[9][11] = newTile(DefualtTile)
	b.Tiles[9][12] = newTile(DefualtTile)
	b.Tiles[9][13] = newTile(TrippleLetter)
	b.Tiles[9][14] = newTile(DefualtTile)

	// Row 11
	b.Tiles[10][0] = newTile(DefualtTile)
	b.Tiles[10][1] = newTile(DefualtTile)
	b.Tiles[10][2] = newTile(DefualtTile)
	b.Tiles[10][3] = newTile(DefualtTile)
	b.Tiles[10][4] = newTile(DoubleWord)
	b.Tiles[10][5] = newTile(DefualtTile)
	b.Tiles[10][6] = newTile(DefualtTile)
	b.Tiles[10][7] = newTile(DefualtTile)
	b.Tiles[10][8] = newTile(DefualtTile)
	b.Tiles[10][9] = newTile(DefualtTile)
	b.Tiles[10][10] = newTile(DoubleWord)
	b.Tiles[10][11] = newTile(DefualtTile)
	b.Tiles[10][12] = newTile(DefualtTile)
	b.Tiles[10][13] = newTile(DefualtTile)
	b.Tiles[10][14] = newTile(DefualtTile)

	// Row 12
	b.Tiles[11][0] = newTile(DoubleLetter)
	b.Tiles[11][1] = newTile(DefualtTile)
	b.Tiles[11][2] = newTile(DefualtTile)
	b.Tiles[11][3] = newTile(DoubleWord)
	b.Tiles[11][4] = newTile(DefualtTile)
	b.Tiles[11][5] = newTile(DefualtTile)
	b.Tiles[11][6] = newTile(DefualtTile)
	b.Tiles[11][7] = newTile(DoubleLetter)
	b.Tiles[11][8] = newTile(DefualtTile)
	b.Tiles[11][9] = newTile(DefualtTile)
	b.Tiles[11][10] = newTile(DefualtTile)
	b.Tiles[11][11] = newTile(DoubleWord)
	b.Tiles[11][12] = newTile(DefualtTile)
	b.Tiles[11][13] = newTile(DefualtTile)
	b.Tiles[11][14] = newTile(DoubleLetter)

	// Row 13
	b.Tiles[12][0] = newTile(DefualtTile)
	b.Tiles[12][1] = newTile(DefualtTile)
	b.Tiles[12][2] = newTile(DoubleWord)
	b.Tiles[12][3] = newTile(DefualtTile)
	b.Tiles[12][4] = newTile(DefualtTile)
	b.Tiles[12][5] = newTile(DefualtTile)
	b.Tiles[12][6] = newTile(DoubleLetter)
	b.Tiles[12][7] = newTile(DefualtTile)
	b.Tiles[12][8] = newTile(DoubleLetter)
	b.Tiles[12][9] = newTile(DefualtTile)
	b.Tiles[12][10] = newTile(DefualtTile)
	b.Tiles[12][11] = newTile(DefualtTile)
	b.Tiles[12][12] = newTile(DoubleWord)
	b.Tiles[12][13] = newTile(DefualtTile)
	b.Tiles[12][14] = newTile(DefualtTile)

	// Row 14
	b.Tiles[13][0] = newTile(DefualtTile)
	b.Tiles[13][1] = newTile(DoubleWord)
	b.Tiles[13][2] = newTile(DefualtTile)
	b.Tiles[13][3] = newTile(DefualtTile)
	b.Tiles[13][4] = newTile(DefualtTile)
	b.Tiles[13][5] = newTile(TrippleLetter)
	b.Tiles[13][6] = newTile(DefualtTile)
	b.Tiles[13][7] = newTile(DefualtTile)
	b.Tiles[13][8] = newTile(DefualtTile)
	b.Tiles[13][9] = newTile(TrippleLetter)
	b.Tiles[13][10] = newTile(DefualtTile)
	b.Tiles[13][11] = newTile(DefualtTile)
	b.Tiles[13][12] = newTile(DefualtTile)
	b.Tiles[13][13] = newTile(DoubleWord)
	b.Tiles[13][14] = newTile(DefualtTile)

	// Row 15
	b.Tiles[14][0] = newTile(TrippleWord)
	b.Tiles[14][1] = newTile(DefualtTile)
	b.Tiles[14][2] = newTile(DefualtTile)
	b.Tiles[14][3] = newTile(DoubleLetter)
	b.Tiles[14][4] = newTile(DefualtTile)
	b.Tiles[14][5] = newTile(DefualtTile)
	b.Tiles[14][6] = newTile(DefualtTile)
	b.Tiles[14][7] = newTile(TrippleWord)
	b.Tiles[14][8] = newTile(DefualtTile)
	b.Tiles[14][9] = newTile(DefualtTile)
	b.Tiles[14][10] = newTile(DefualtTile)
	b.Tiles[14][11] = newTile(DoubleLetter)
	b.Tiles[14][12] = newTile(DefualtTile)
	b.Tiles[14][13] = newTile(DefualtTile)
	b.Tiles[14][14] = newTile(TrippleWord)
	return
}
