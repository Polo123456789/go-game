package trender

// Because all "pixels" are just ascii characters, to make them look more
// squared I will use 3 characters for each pixel.

const VisualPixelWidth = 3

type Pixel interface {
	ToAnsiEscapeCode() string
	SetGraphicsMode(mode GraphicsMode)
	SetContent(c rune)
}
