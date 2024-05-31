package trender

import (
	"io"
	"os"
	"strconv"
)

const (
	AnsiClearScreen = "\x1b[2J\x1b[H"
)

type canvasPixel struct {
	Pixel   Pixel
	Changed bool
}

type Canvas struct {
	// Indexed by [y][x]
	pixels    [][]canvasPixel
	width     int
	height    int
	buffer    []byte
	ansiCache map[uint64]string
	writer    io.Writer
}

func (c *Canvas) Width() int {
	return c.width
}

func (c *Canvas) Height() int {
	return c.height
}

func NewCanvas(width, height int, defaultPixel Pixel) *Canvas {
	pixels := make([][]canvasPixel, height)
	for y := 0; y < height; y++ {
		pixels[y] = make([]canvasPixel, width)
		for x := 0; x < width; x++ {
			pixels[y][x] = canvasPixel{Pixel: defaultPixel, Changed: true}
		}
	}
	return &Canvas{
		width:     width,
		height:    height,
		pixels:    pixels,
		writer:    os.Stdout,
		buffer:    make([]byte, 0, defaultPixel.MaxPossibleSize()*width*height),
		ansiCache: make(map[uint64]string),
	}
}

func (c *Canvas) SetWriter(writer io.Writer) {
	c.writer = writer
}

func (c *Canvas) SetPixel(x, y int, pixel Pixel) {
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
	c.bufferAppend("\x1b[")
	c.bufferAppend(strconv.Itoa(y))
	c.bufferAppend(";")
	c.bufferAppend(strconv.Itoa(x))
	c.bufferAppend("H")
}

// func SetCursorPosition(x, y int) string {
// 	return "\x1b[" + strconv.Itoa(y) + ";" + strconv.Itoa(x) + "H"
// }

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
