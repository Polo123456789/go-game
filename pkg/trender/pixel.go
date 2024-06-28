package trender

const PixelWidth = 2

// Pixel should be 2 characters wide and 1 character tall
type Pixel interface {
	ToAnsiEscapeCode() string
	SetContent(c rune)
	MaxPossibleSize() int
	HashKey() uint64
}
