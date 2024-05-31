package pixels

import (
	"strconv"
)

type ColorID uint8

type Pixel struct {
	Foreground ColorID
	Background ColorID
	Content    rune
}

func NewPixel(foreground, background ColorID, content rune) *Pixel {
	return &Pixel{
		Foreground: foreground,
		Background: background,
		Content:    content,
	}
}

func (p *Pixel) ToAnsiEscapeCode() string {
	out := ""
	out += "\x1b[38;5;" + strconv.Itoa(int(p.Foreground)) + "m"
	out += "\x1b[48;5;" + strconv.Itoa(int(p.Background)) + "m"
	out += string(p.Content)
	return out
}

func (p *Pixel) SetContent(c rune) {
	p.Content = c
}

func (p *Pixel) MaxPossibleSize() int {
	// TODO: Implement
	panic("not implemented")
}

func (p *Pixel) HashKey() uint64 {
	// TODO: Implement
	panic("not implemented")
}
