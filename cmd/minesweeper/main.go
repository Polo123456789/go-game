package main

import (
	"flag"
	"fmt"

	"github.com/Polo123456789/go-game/pkg/input"

	board "github.com/Polo123456789/go-game/pkg/mines-board"
	boardUI "github.com/Polo123456789/go-game/pkg/mines-ui"
)

const (
	DefaultRows  = 8
	DefaultCols  = 8
	DefaultMines = 10
)

var (
	rows  = flag.Int("rows", DefaultRows, "Number of rows")
	cols  = flag.Int("cols", DefaultCols, "Number of columns")
	mines = flag.Int("mines", DefaultMines, "Number of mines")
)

func main() {
	flag.Parse()
	input.UnbufferStdin()
	defer input.RestoreStdin()

	b := board.NewBoard(*rows, *cols, *mines)
	ui := boardUI.NewBoardUI(&b, boardUI.Cursor{X: 0, Y: 0})
	matchTree := input.NewMatchTree(boardUI.InputTree)

	for {
		boardUI.ClearScreen()
		ui.Draw()
		fmt.Println("\nMove with wasd of hjkl, (c)lear, (f)lag, (m)ark,")
		fmt.Println("\t(r)eset tile state, (q)uit")
		userInput := input.NoMatch
		modifier := input.DefaultNumericModifier
		for userInput == input.NoMatch {
			matchTree.MatchOrReset(input.Get())
			if matchTree.CurrentResult() != input.NoMatch {
				userInput = matchTree.CurrentResult()
				modifier = matchTree.NumericModifier()
				matchTree.Reset()
			}
		}

		var result board.MoveResult
		switch {
		case userInput >= boardUI.MoveUp && userInput <= boardUI.MoveAllRight:
			ui.MoveCursor(userInput, modifier)
			result = board.MoveResultSuccessful
		case userInput == boardUI.Clear:
			result = ui.MakeMoveAtCursor(board.PlayerMarkedCleared)
		case userInput == boardUI.Flag:
			result = ui.MakeMoveAtCursor(board.PlayerMarkedDoubtful)
		case userInput == boardUI.Mark:
			result = ui.MakeMoveAtCursor(board.PlayerMarkedMined)
		case userInput == boardUI.ClearState:
			result = ui.MakeMoveAtCursor(board.PlayerClearedState)
		case userInput == boardUI.Quit:
			return
		}

		if result&board.MoveResultDeath != 0 {
			boardUI.ClearScreen()
			ui.Draw()
			fmt.Println("You died!")
			return
		} else if b.GameOver() {
			boardUI.ClearScreen()
			ui.Draw()
			fmt.Println("You won!")
			return
		}
	}
}
