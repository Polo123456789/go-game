package trender

type Pixel interface {
	// ToAnsiEscapeCode should be deterministic, as it might be used in the
	// future to generate a cache, or to compare pixels in a double buffer
	ToAnsiEscapeCode() string

	SetGraphicsMode(mode GraphicsMode)
	SetContent(c rune)
}
