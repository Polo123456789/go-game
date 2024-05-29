package pixels

import (
	"github.com/Polo123456789/go-game/pkg/trender"
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
	Reset
)

type Pixel struct {
	Foreground Color
	Background Color
	Content    rune
	trender.GraphicsMode
}

func NewPixel(foreground, background Color, content rune) *Pixel {
	return &Pixel{
		Foreground:   foreground,
		Background:   background,
		Content:      content,
		GraphicsMode: trender.Reset,
	}
}

func (p *Pixel) ToAnsiEscapeCode() string {
	f := strconv.Itoa(int(30 + p.Foreground))
	b := strconv.Itoa(int(40 + p.Background))
	m := strconv.Itoa(int(p.GraphicsMode))
	return "\x1b[" + m + ";" + f + ";" + b + "m" + string(p.Content)
}

func (p *Pixel) SetGraphicsMode(mode trender.GraphicsMode) {
	p.GraphicsMode = mode
}

func (p *Pixel) SetContent(c rune) {
	p.Content = c
}
