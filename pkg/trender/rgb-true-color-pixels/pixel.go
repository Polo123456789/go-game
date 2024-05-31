package pixels

import (
	"fmt"
)

type RGB struct {
	R, G, B uint8
}

type Pixel struct {
	Foreground RGB
	Background RGB
	Content    rune
}

func NewPixel(foreground, background RGB, content rune) *Pixel {
	return &Pixel{
		Foreground: foreground,
		Background: background,
		Content:    content,
	}
}

func (p *Pixel) ToAnsiEscapeCode() string {
	out := fmt.Sprintf("\x1b[38;2;%d;%d;%dm", p.Foreground.R, p.Foreground.G, p.Foreground.B)
	out += fmt.Sprintf("\x1b[48;2;%d;%d;%dm", p.Background.R, p.Background.G, p.Background.B)
	return out + string(p.Content)
}

func (p *Pixel) SetContent(c rune) {
	p.Content = c
}

func (p *Pixel) MaxPossibleSize() int {
	const longestPossible = "\x1b[48;2;255;255;255m"
	return len(longestPossible)*2 + 1
}

func (p *Pixel) HashKey() uint64 {
	return uint64(p.Foreground.R) |
		uint64(p.Foreground.G)<<8 |
		uint64(p.Foreground.B)<<16 |
		uint64(p.Background.R)<<24 |
		uint64(p.Background.G)<<32 |
		uint64(p.Background.B)<<40 |
		uint64(p.Content)<<48
}
