package trender

import (
	"io"
	"os"
)

type canvasPixel struct {
	Pixel   Pixel
	Changed bool
}

type Canvas struct {
	// Indexed by [y][x]
	pixels [][]canvasPixel
	width  int
	height int
	buffer string
	writer io.Writer
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
		width:  width,
		height: height,
		pixels: pixels,
		writer: os.Stdout,
	}
}

func (c *Canvas) SetWriter(writer io.Writer) {
	c.writer = writer
}

func (c *Canvas) SetPixel(x, y int, pixel Pixel) {
	c.pixels[y][x] = canvasPixel{Pixel: pixel, Changed: true}
}

// PreRender fills the canvas buffer with all the pixels that have changed
// since the last render.
//
// Its mean to be used in a separate goroutine to avoid blocking the main
// thread. You can pass in nil to the channel if you wish to run it
// syncronously.
//
// Example:
//
// ```go
//
//	canvas := trender.NewCanvas(...)
//	// Setup the pixels
//	preRenderChannel := make(chan bool)
//	go canvas.PreRender(preRenderChannel)
//
//	// Other things that need to be done
//
//	<-preRenderChannel
//	canvas.Render()
//
// ```
func (c *Canvas) PreRender(channel chan bool) {
	for y, row := range c.pixels {
		for x, pixel := range row {
			if pixel.Changed {
				c.buffer += SetCursorPosition(x, y)
				c.buffer += pixel.Pixel.ToAnsiEscapeCode()
				c.pixels[y][x].Changed = false
			}
		}
	}
	if channel != nil {
		channel <- true
	}
}

// Render writes the canvas buffer to the writer.
func (c *Canvas) Render() {
	if c.buffer == "" {
		c.PreRender(nil)
	}
	c.writer.Write([]byte(c.buffer))
	c.buffer = ""
}

// FullPrerender does the same thing as PreRender, but ignores if there was a
// change to a pixel or not
func (c *Canvas) FullPrerender(channel chan bool) {
	c.buffer = SetCursorPosition(0, 0)
	for y, row := range c.pixels {
		for x, pixel := range row {
			c.buffer += pixel.Pixel.ToAnsiEscapeCode()
			c.pixels[y][x].Changed = false
		}
		c.buffer += SetCursorPosition(0, y+1)
	}
	if channel != nil {
		channel <- true
	}
}

// FullRender does the same thing as Render, but ignores if there was a change
// to a pixel or not
func (c *Canvas) FullRender() {
	if c.buffer == "" {
		c.FullPrerender(nil)
	}
	c.writer.Write([]byte(c.buffer))
	c.buffer = ""
}
