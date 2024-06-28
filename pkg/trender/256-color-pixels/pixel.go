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
	out += string(p.Content) + string(p.Content)
	return out
}

func (p *Pixel) SetContent(c rune) {
	p.Content = c
}

func (p *Pixel) MaxPossibleSize() int {
	const longestPossible = "\x1b[48;5;255m"
	return len(longestPossible)*2 + 2
}

func (p *Pixel) HashKey() uint64 {
	return uint64(p.Foreground) |
		uint64(p.Background)<<8 |
		uint64(p.Content)<<16
}
