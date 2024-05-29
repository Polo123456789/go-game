package pixels

import (
	"github.com/Polo123456789/go-game/pkg/trender"
	"strconv"
)

type ColorID uint16

const DefaultColor ColorID = 256 // Valid ids are 0-255

type Pixel struct {
	Foreground ColorID
	Background ColorID
	Content    rune
	trender.GraphicsMode
}

func NewPixel(foreground, background ColorID, content rune) *Pixel {
	return &Pixel{
		Foreground:   foreground,
		Background:   background,
		Content:      content,
		GraphicsMode: trender.Reset,
	}
}

func (p *Pixel) ToAnsiEscapeCode() string {
	out := ""
	out += "\x1b[38;5;" + strconv.Itoa(int(p.Foreground)) + "m"
	out += "\x1b[48;5;" + strconv.Itoa(int(p.Background)) + "m"
	content := string(p.Content)

	return out + " " + content + " "
}

func (p *Pixel) SetContent(c rune) {
	p.Content = c
}

func (p *Pixel) SetGraphicsMode(_ trender.GraphicsMode) {}
