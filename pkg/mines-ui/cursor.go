package ui

import (
	"github.com/Polo123456789/go-game/pkg/input"
)

type Cursor struct {
	X, Y int
}

func (ui *BoardUi) MoveCursor(move input.TreeResult) {
	switch move {
	case MoveUp:
		if ui.Cursor.Y > 0 {
			ui.Cursor.Y--
		}
	case MoveDown:
		if ui.Cursor.Y < ui.Board.Height-1 {
			ui.Cursor.Y++
		}
	case MoveLeft:
		if ui.Cursor.X > 0 {
			ui.Cursor.X--
		}
	case MoveRight:
		if ui.Cursor.X < ui.Board.Width-1 {
			ui.Cursor.X++
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
