package pixels

import (
	"strconv"
)

type Color uint8

const (
	Black Color = iota
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
	Default
	ResetColor
)

type Pixel struct {
	Foreground Color
	Background Color
	Content    rune
	GraphicsMode
}

func NewPixel(
	foreground Color,
	background Color,
	graphicsMode GraphicsMode,
	content rune,
) *Pixel {
	return &Pixel{
		Foreground:   foreground,
		Background:   background,
		Content:      content,
		GraphicsMode: graphicsMode,
	}
}

func (p *Pixel) ToAnsiEscapeCode() string {
	f := strconv.Itoa(int(30 + p.Foreground))
	b := strconv.Itoa(int(40 + p.Background))
	m := strconv.Itoa(int(p.GraphicsMode))
	return "\x1b[" + m + ";" + f + ";" + b + "m" + string(p.Content)
}

func (p *Pixel) SetContent(c rune) {
	p.Content = c
}

func (p *Pixel) MaxPossibleSize() int {
	const longestPossible = "\x1b[9;30;40m"
	return len(longestPossible) + 1
}

func (p *Pixel) HashKey() uint64 {
	return uint64(p.Foreground) |
		uint64(p.Background)<<8 |
		uint64(p.GraphicsMode)<<16 |
		uint64(p.Content)<<24
}
