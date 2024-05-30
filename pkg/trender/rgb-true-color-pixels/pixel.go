package pixels

import (
	"fmt"
)

type RGB struct {
	R, G, B uint16
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
