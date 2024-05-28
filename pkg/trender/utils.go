package trender

import (
	"strconv"
)

const (
	AnsiClearScreen = "\x1b[2J\x1b[H"
)

func SetCursorPosition(x, y int) string {
	return "\x1b[" + strconv.Itoa(y) + ";" + strconv.Itoa(x) + "H"
}
