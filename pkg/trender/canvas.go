package trender

import (
	"io"
	"os"
	"strconv"
)

const (
	AnsiClearScreen = "\x1b[2J\x1b[H"
)

func ClearScreen() {
	os.Stdout.Write([]byte(AnsiClearScreen))
}

type canvasPixel struct {
	Pixel   Pixel
	Changed bool
}

type Canvas struct {
	// Indexed by [y][x]
	pixels          [][]canvasPixel
	width           int
	height          int
	buffer          []byte
	ansiCache       map[uint64]string
	cursorPositions [][][]byte
	writer          io.Writer
}

func (c *Canvas) Width() int {
	return c.width
}

func (c *Canvas) Height() int {
	return c.height
}

func NewCanvas(width, height int, defaultPixel Pixel) *Canvas {
	pixels := make([][]canvasPixel, height)
	cursorPositions := make([][][]byte, height)
	for y := 0; y < height; y++ {
		pixels[y] = make([]canvasPixel, width)
		cursorPositions[y] = make([][]byte, width)
		for x := 0; x < width; x++ {
			pixels[y][x] = canvasPixel{Pixel: defaultPixel, Changed: true}
			cursorPositions[y][x] = []byte("\x1b[" + strconv.Itoa(y+1) + ";" + strconv.Itoa(x+1) + "H")
		}
	}
	return &Canvas{
		width:           width,
		height:          height,
		pixels:          pixels,
		writer:          os.Stdout,
		buffer:          make([]byte, 0, defaultPixel.MaxPossibleSize()*width*height),
		cursorPositions: cursorPositions,
		ansiCache:       make(map[uint64]string),
	}
}

func (c *Canvas) SetWriter(writer io.Writer) {
	c.writer = writer
}

func (c *Canvas) SetPixel(x, y int, pixel Pixel) {
	if x < 0 {
		x = 0
	}
	if x >= c.width {
		x = c.width - 1
	}
	if y < 0 {
		y = 0
	}
	if y >= c.height {
		y = c.height - 1
	}
	c.pixels[y][x] = canvasPixel{Pixel: pixel, Changed: true}
}

func (c *Canvas) bufferAppend(s string) {
	c.buffer = append(c.buffer, s...)
}

func (c *Canvas) bufferReset() {
	c.buffer = c.buffer[:0]
}

func (c *Canvas) getAnsiRepresentation(p Pixel) string {
	cached, ok := c.ansiCache[p.HashKey()]
	if ok {
		return cached
	}

	ansi := p.ToAnsiEscapeCode()
	c.ansiCache[p.HashKey()] = ansi
	return ansi
}

func (c *Canvas) setCursorPosition(x, y int) {
	c.buffer = append(c.buffer, c.cursorPositions[y][x]...)
}

func (c *Canvas) RenderChanged() {
	c.bufferReset()
	for y, row := range c.pixels {
		for x, pixel := range row {
			if pixel.Changed {
				c.setCursorPosition(x, y)
				c.bufferAppend(c.getAnsiRepresentation(pixel.Pixel))
				c.pixels[y][x].Changed = false
			}
		}
	}
	c.writer.Write(c.buffer)
}

func (c *Canvas) RenderFull() {
	c.bufferReset()
	c.setCursorPosition(0, 0)
	for y, row := range c.pixels {
		for x, pixel := range row {
			c.bufferAppend(c.getAnsiRepresentation(pixel.Pixel))
			c.pixels[y][x].Changed = false
		}
		c.setCursorPosition(0, y+1)
	}
	c.writer.Write(c.buffer)
}

func (c *Canvas) Clear(p Pixel) {
	for y, row := range c.pixels {
		for x := range row {
			if c.pixels[y][x].Pixel.HashKey() != p.HashKey() {
				c.SetPixel(x, y, p)
			}
		}
	}
}

func (c *Canvas) DrawRect(r Rect, p Pixel) {
	panic("Not Implemented")
}

func (c *Canvas) DrawRectFill(r Rect, p Pixel) {
	panic("Not Implemented")
}
