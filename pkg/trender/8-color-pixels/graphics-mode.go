package pixels

type GraphicsMode uint8

const (
	Reset GraphicsMode = iota
	Bold
	Dim
	Italic
	Underline
	Blink
	Inverse
	Hidden
	Strikethrough
)
