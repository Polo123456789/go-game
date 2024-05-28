package main

import (
	"fmt"
)

type BoardUI struct {
	Board  *Board
	Cursor Cursor
}

func NewBoardUI(board *Board, startingPosition Cursor) BoardUI {
	return BoardUI{
		Board:  board,
		Cursor: startingPosition,
	}
}

func (ui *BoardUI) MakeMoveAtCursor(move PlayerMove) MoveResult {
	return ui.Board.MakeMove(ui.Cursor.X, ui.Cursor.Y, move)
}

func drawTile(tile Tile) {
	switch tile.PlayerView {
	case UndiscoveredTile:
		fmt.Print(" # ")
	case ClearTile:
		if tile.Value != MinedTile {
			fmt.Printf(" %v ", tile.Value)
		} else {
			fmt.Print(" M ")
		}
	case MarkedMinedTile:
		fmt.Print(" M ")
	case MarkedDoubtfulTile:
		fmt.Print(" ? ")
	default:
		// This should never happen
		panic(fmt.Sprintf("Invalid PlayerView %d", tile.PlayerView))
	}
}

func (ui *BoardUI) Draw() {
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

func tileGraphicsMode(tile Tile) {
	switch tile.PlayerView {
	case ClearTile:
		if tile.Value != MinedTile {
			fmt.Print("\x1b[30;107m")
		} else {
			fmt.Print("\x1b[30;41m")
		}
	case MarkedMinedTile:
		fmt.Print("\x1b[48;5;208m")
		fmt.Print("\x1b[1;30m")
	case MarkedDoubtfulTile:
		fmt.Print("\x1b[48;5;226m")
		fmt.Print("\x1b[1;30m")
	case UndiscoveredTile:
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
