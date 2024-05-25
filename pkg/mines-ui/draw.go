package ui

import (
	"fmt"
	"github.com/Polo123456789/go-game/pkg/mines-board"
)

type BoardUi struct {
	Board  *board.Board
	Cursor Cursor
}

func NewBoardUi(board *board.Board, startingPosition Cursor) BoardUi {
	return BoardUi{
		Board:  board,
		Cursor: startingPosition,
	}
}

func (ui *BoardUi) MakeMoveAtCursor(move board.PlayerMove) board.MoveResult {
	return ui.Board.MakeMove(ui.Cursor.X, ui.Cursor.Y, move)
}

func drawTile(tile board.Tile) {
	switch tile.PlayerView {
	case board.UndiscoveredTile:
		fmt.Print(" # ")
	case board.ClearTile:
		if tile.Value != board.MinedTile {
			fmt.Printf(" %v ", tile.Value)
		} else {
			fmt.Print(" M ")
		}
	case board.MarkedMinedTile:
		fmt.Print(" M ")
	case board.MarkedDoubtfulTile:
		fmt.Print(" ? ")
	default:
		// This should never happen
		panic(fmt.Sprintf("Invalid PlayerView %d", tile.PlayerView))
	}
}

func (ui *BoardUi) Draw() {
	for y := 0; y < ui.Board.Height; y++ {
		for x := 0; x < ui.Board.Width; x++ {
			tile := ui.Board.Tiles[x][y]
			if ui.Cursor.X == x && ui.Cursor.Y == y {
				cursorGraphicsMode()
				drawTile(tile)
				resetGraphicsMode()
			} else {
				tileGraphicsMode(tile)
				drawTile(tile)
				resetGraphicsMode()
			}
		}
		fmt.Print("\n")
	}
}

// {{{ Graphics modes
func resetGraphicsMode() {
	fmt.Print("\x1b[0m")
}

func cursorGraphicsMode() {
	fmt.Print("\x1b[5;30;47m")
}

func tileGraphicsMode(tile board.Tile) {
	switch tile.PlayerView {
	case board.ClearTile:
		if tile.Value != board.MinedTile {
			fmt.Print("\x1b[30;107m")
		} else {
			fmt.Print("\x1b[30;41m")
		}
	case board.MarkedMinedTile:
		fmt.Print("\x1b[48;5;208m")
		fmt.Print("\x1b[1;30m")
	case board.MarkedDoubtfulTile:
		fmt.Print("\x1b[48;5;226m")
		fmt.Print("\x1b[1;30m")
	case board.UndiscoveredTile:
		fmt.Print("\x1b[48;5;240m")
	default:
		// This should never happen
		panic(fmt.Sprintf("Invalid PlayerView %d", tile.PlayerView))
	}
}

func ClearScreen() {
	fmt.Print("\x1b[2J")
	fmt.Print("\x1b[H")
}

// }}}
