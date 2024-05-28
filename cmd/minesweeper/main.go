package main

import (
	"flag"
	"fmt"

	"github.com/Polo123456789/go-game/pkg/input"

	board "github.com/Polo123456789/go-game/pkg/mines-board"
	bui "github.com/Polo123456789/go-game/pkg/mines-ui"
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
	ui := bui.NewBoardUi(&b, bui.Cursor{X: 0, Y: 0})
	matchTree := input.NewMatchTree(bui.InputTree)

	for {
		bui.ClearScreen()
		ui.Draw()
		fmt.Println("\nMove with wasd of hjkl, (c)lear, (f)lag, (m)ark,")
		fmt.Println("\t(r)eset tile state, (q)uit")
		// TODO: move to using inputTree
		// input := bui.TranslateInput(input.Get())
		userInput := input.NoMatch
		for userInput == input.NoMatch {
			matchTree.MatchOrReset(input.Get())
			if matchTree.CurrentResult() != input.NoMatch {
				userInput = matchTree.CurrentResult()
				matchTree.Reset()
			}
		}

		var result board.MoveResult
		switch userInput {
		case bui.MoveUp, bui.MoveLeft, bui.MoveDown, bui.MoveRight,
			bui.MoveAllUp, bui.MoveAllLeft, bui.MoveAllDown, bui.MoveAllRight:
			ui.MoveCursor(userInput)
			result = board.MoveResultSuccessful
		case bui.Clear:
			result = ui.MakeMoveAtCursor(board.PlayerMarkedCleared)
		case bui.Flag:
			result = ui.MakeMoveAtCursor(board.PlayerMarkedDoubtful)
		case bui.Mark:
			result = ui.MakeMoveAtCursor(board.PlayerMarkedMined)
		case bui.ClearState:
			result = ui.MakeMoveAtCursor(board.PlayerClearedState)
		case bui.Quit:
			return
		}

		if result&board.MoveResultDeath != 0 {
			bui.ClearScreen()
			ui.Draw()
			fmt.Println("You died!")
			return
		} else if b.GameOver() {
			bui.ClearScreen()
			ui.Draw()
			fmt.Println("You won!")
			return
		}
	}
}
