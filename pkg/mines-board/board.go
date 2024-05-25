package board

import (
	"fmt"
	"math"
	"math/rand"
)

type Tile struct {
	Value      int
	PlayerView PlayerViewState
}

type Board struct {
	Tiles         [][]Tile
	Width         int
	Height        int
	firstMoveMade bool
	noMines       int
}

// {{{Constants
type MoveResult int

const (
	MoveResultSuccessful MoveResult = 1 << iota
	MoveResultDeath
	MoveResultDidNothing
)

/**
 * if tile > 0, then tile represents the number of adjacent mines
 * if tile < 0, then the tile is mined
 */
const (
	EmptyTile = 0
	MinedTile = -1
)

type PlayerViewState int

const (
	UndiscoveredTile PlayerViewState = iota
	ClearTile
	MarkedMinedTile
	MarkedDoubtfulTile
)

type PlayerMove int

const (
	PlayerMarkedCleared PlayerMove = iota
	PlayerMarkedDoubtful
	PlayerMarkedMined
	PlayerClearedState
)

// }}}

// {{{ Board Generation
func newTileArray(width, height int) [][]Tile {
	array := make([][]Tile, width)
	for x := range array {
		array[x] = make([]Tile, height)
		for y := range array[x] {
			array[x][y] = Tile{
				Value:      EmptyTile,
				PlayerView: UndiscoveredTile,
			}
		}
	}
	return array
}

func NewBoard(width, height, noMines int) Board {
	board := Board{
		Tiles:         newTileArray(width, height),
		Width:         width,
		Height:        height,
		firstMoveMade: false,
		noMines:       noMines,
	}

	return board
}

func placeRandomMines(b *Board, excludeX, excludeY int) {
	type PossibleLocation struct {
		x, y int
	}

	var possibleLocations []PossibleLocation
	for x := 0; x < b.Width; x++ {
		for y := 0; y < b.Height; y++ {
			if math.Abs(float64(x-excludeX)) < 2 && math.Abs(float64(y-excludeY)) < 2 {
				continue
			}
			possibleLocations = append(possibleLocations, PossibleLocation{x, y})
		}
	}

	rand.Shuffle(len(possibleLocations), func(i, j int) {
		possibleLocations[i], possibleLocations[j] =
			possibleLocations[j], possibleLocations[i]
	})

	locationsToMine := possibleLocations[:b.noMines]
	for _, loc := range locationsToMine {
		b.Tiles[loc.x][loc.y].Value = MinedTile
	}
}

func calculateTileValues(b *Board) {
	minesAt := func(x, y int) int {
		if x < 0 || x >= b.Width {
			return 0
		}
		if y < 0 || y >= b.Height {
			return 0
		}
		if b.Tiles[x][y].Value == MinedTile {
			return 1
		}
		return 0
	}

	for x := 0; x < b.Width; x++ {
		for y := 0; y < b.Height; y++ {
			if b.Tiles[x][y].Value == MinedTile {
				continue
			}

			b.Tiles[x][y].Value =
				minesAt(x-1, y-1) + minesAt(x, y-1) + minesAt(x+1, y-1) +
					minesAt(x-1, y) /*Ignore current*/ + minesAt(x+1, y) +
					minesAt(x-1, y+1) + minesAt(x, y+1) + minesAt(x+1, y+1)
		}
	}
}

func populateBoard(b *Board, excludeX, excludeY int) {
	placeRandomMines(b, excludeX, excludeY)
	calculateTileValues(b)
}

// }}}

// {{{ Board Moves
func (b *Board) movePlayerClearedState(x, y int) MoveResult {
	tile := &b.Tiles[x][y]

	switch tile.PlayerView {
	case UndiscoveredTile, ClearTile:
		return MoveResultDidNothing
	case MarkedMinedTile, MarkedDoubtfulTile:
		tile.PlayerView = UndiscoveredTile
		return MoveResultSuccessful
	default:
		// This should never happen
		panic(fmt.Sprintf("Invalid tile state: %d", tile.PlayerView))
	}
}

func (b *Board) floodClearZeroes(x, y int) {
	if x < 0 || x >= b.Width || y < 0 || y >= b.Height {
		return
	}
	tile := &b.Tiles[x][y]

	if tile.PlayerView != UndiscoveredTile {
		return
	}

	tile.PlayerView = ClearTile

	if tile.Value == 0 {
		b.floodClearZeroes(x-1, y-1)
		b.floodClearZeroes(x, y-1)
		b.floodClearZeroes(x+1, y-1)
		b.floodClearZeroes(x-1, y)
		b.floodClearZeroes(x+1, y)
		b.floodClearZeroes(x-1, y+1)
		b.floodClearZeroes(x, y+1)
		b.floodClearZeroes(x+1, y+1)
	}
}

func (b *Board) movePlayerMarkedCleared(x, y int) MoveResult {
	tile := &b.Tiles[x][y]

	if tile.PlayerView != UndiscoveredTile {
		return MoveResultDidNothing
	}

	if tile.Value == MinedTile {
		for x := 0; x < b.Width; x++ {
			for y := 0; y < b.Height; y++ {
				b.Tiles[x][y].PlayerView = ClearTile
			}
		}
		return MoveResultDeath
	}

	if tile.Value == 0 {
		b.floodClearZeroes(x, y)
	} else {
		tile.PlayerView = ClearTile
	}

	return MoveResultSuccessful
}

func (b *Board) movePlayerMarkedDoubtful(x, y int) MoveResult {
	tile := &b.Tiles[x][y]

	switch tile.PlayerView {
	case UndiscoveredTile, MarkedMinedTile:
		tile.PlayerView = MarkedDoubtfulTile
		return MoveResultSuccessful
	case MarkedDoubtfulTile:
		tile.PlayerView = UndiscoveredTile
		return MoveResultSuccessful
	case ClearTile:
		return MoveResultDidNothing
	default:
		// This should never happen
		panic(fmt.Sprintf("Invalid tile state: %d", tile.PlayerView))
	}
}

func (b *Board) movePlayerMarkedMined(x, y int) MoveResult {
	tile := &b.Tiles[x][y]

	switch tile.PlayerView {
	case UndiscoveredTile, MarkedDoubtfulTile:
		tile.PlayerView = MarkedMinedTile
		return MoveResultSuccessful
	case MarkedMinedTile:
		tile.PlayerView = UndiscoveredTile
		return MoveResultSuccessful
	case ClearTile:
		return MoveResultDidNothing
	default:
		// This should never happen
		panic(fmt.Sprintf("Invalid tile state: %d", tile.PlayerView))
	}
}

func (b *Board) MakeMove(x, y int, move PlayerMove) MoveResult {
	if !b.firstMoveMade {
		populateBoard(b, x, y)
		b.firstMoveMade = true
	}

	switch move {
	case PlayerMarkedCleared:
		return b.movePlayerMarkedCleared(x, y)
	case PlayerMarkedDoubtful:
		return b.movePlayerMarkedDoubtful(x, y)
	case PlayerMarkedMined:
		return b.movePlayerMarkedMined(x, y)
	case PlayerClearedState:
		return b.movePlayerClearedState(x, y)
	default:
		// This should never happen
		panic(fmt.Sprintf("Invalid move: %d", move))
	}
}

// }}}

func (b *Board) GameOver() bool {
	noMarkedTiles := 0
	for x := 0; x < b.Width; x++ {
		for y := 0; y < b.Height; y++ {
			switch b.Tiles[x][y].PlayerView {
			case ClearTile:
				continue
			case MarkedMinedTile:
				noMarkedTiles++
			default:
				return false
			}
		}
	}
	return noMarkedTiles == b.noMines
}
