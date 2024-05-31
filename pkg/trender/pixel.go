package trender

type Pixel interface {
	ToAnsiEscapeCode() string
	SetContent(c rune)
	MaxPossibleSize() int
	HashKey() uint64
}
