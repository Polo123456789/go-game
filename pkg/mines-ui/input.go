package ui

import (
	"os"
	"os/exec"
)

// {{{ Stdin
func UnbufferStdin() {
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
}

func RestoreStdin() {
	defer exec.Command("stty", "-F", "/dev/tty", "echo").Run()
}

func Input() byte {
	b := [1]byte{}
	os.Stdin.Read(b[:])
	return b[0]
}

// }}}

// {{{ Movement
type Movement int

const (
	MoveUp Movement = iota
	MoveDown
	MoveLeft
	MoveRight
	Clear
	Flag
	Mark
	Quit
	ClearState
	Invalid
)

func TranslateInput(input byte) Movement {
	switch input {
	case 'w', 'k':
		return MoveUp
	case 's', 'j':
		return MoveDown
	case 'a', 'h':
		return MoveLeft
	case 'd', 'l':
		return MoveRight
	case 'c':
		return Clear
	case 'f':
		return Flag
	case 'm':
		return Mark
	case 'r':
		return ClearState
	case 'q':
		return Quit
	}
	return Invalid
}

// }}}
