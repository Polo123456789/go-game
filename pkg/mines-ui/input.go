package ui

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
