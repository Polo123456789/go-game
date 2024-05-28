package ui

import (
	"github.com/Polo123456789/go-game/pkg/input"
)

type Cursor struct {
	X, Y int
}

func (ui *BoardUI) MoveCursor(move input.TreeResult, modifier int) {
	switch move {
	case MoveUp:
		ui.Cursor.Y -= modifier
		if ui.Cursor.Y < 0 {
			ui.Cursor.Y = 0
		}
	case MoveDown:
		ui.Cursor.Y += modifier
		if ui.Cursor.Y >= ui.Board.Height {
			ui.Cursor.Y = ui.Board.Height - 1
		}
	case MoveLeft:
		ui.Cursor.X -= modifier
		if ui.Cursor.X < 0 {
			ui.Cursor.X = 0
		}
	case MoveRight:
		ui.Cursor.X += modifier
		if ui.Cursor.X >= ui.Board.Width {
			ui.Cursor.X = ui.Board.Width - 1
		}
	case MoveAllUp:
		ui.Cursor.Y = 0
	case MoveAllDown:
		ui.Cursor.Y = ui.Board.Height - 1
	case MoveAllLeft:
		ui.Cursor.X = 0
	case MoveAllRight:
		ui.Cursor.X = ui.Board.Width - 1
	}
}
