package main

import (
	"flag"
	"fmt"

	"github.com/Polo123456789/go-game/pkg/input"
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

	b := NewBoard(*rows, *cols, *mines)
	ui := NewBoardUI(&b, Cursor{X: 0, Y: 0})
	matchTree := input.NewMatchTree(InputTree)

	for {
		ClearScreen()
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

		var result MoveResult
		switch {
		case userInput >= MoveUp && userInput <= MoveAllRight:
			ui.MoveCursor(userInput, modifier)
			result = MoveResultSuccessful
		case userInput == Clear:
			result = ui.MakeMoveAtCursor(PlayerMarkedCleared)
		case userInput == Flag:
			result = ui.MakeMoveAtCursor(PlayerMarkedDoubtful)
		case userInput == Mark:
			result = ui.MakeMoveAtCursor(PlayerMarkedMined)
		case userInput == ClearState:
			result = ui.MakeMoveAtCursor(PlayerClearedState)
		case userInput == Quit:
			return
		}

		if result&MoveResultDeath != 0 {
			ClearScreen()
			ui.Draw()
			fmt.Println("You died!")
			return
		} else if b.GameOver() {
			ClearScreen()
			ui.Draw()
			fmt.Println("You won!")
			return
		}
	}
}
