package trender

type Pixel interface {
	ToAnsiEscapeCode() string
	SetGraphicsMode(mode GraphicsMode)
	SetContent(c rune)
}
