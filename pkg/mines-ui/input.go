package ui

import (
	"github.com/Polo123456789/go-game/pkg/input"
)

const (
	MoveUp input.TreeResult = iota
	MoveDown
	MoveLeft
	MoveRight
	MoveAllUp
	MoveAllDown
	MoveAllLeft
	MoveAllRight
	Clear
	Flag
	Mark
	Quit
	ClearState
)

var InputTree []input.MatchTreeElement = []input.MatchTreeElement{
	// WASD
	{Value: "w", Result: MoveUp},
	{Value: "s", Result: MoveDown},
	{Value: "a", Result: MoveLeft},
	{Value: "d", Result: MoveRight},

	// Vim Movements
	{Value: "k", Result: MoveUp},
	{Value: "j", Result: MoveDown},
	{Value: "h", Result: MoveLeft},
	{Value: "l", Result: MoveRight},
	{Value: "gg", Result: MoveAllUp},
	{Value: "G", Result: MoveAllDown},
	{Value: "0", Result: MoveAllLeft},
	{Value: "$", Result: MoveAllRight},

	// Other
	{Value: "c", Result: Clear},
	{Value: "f", Result: Flag},
	{Value: "m", Result: Mark},
	{Value: "q", Result: Quit},
	{Value: "r", Result: ClearState},
}
