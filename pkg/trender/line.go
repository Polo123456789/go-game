package trender

import (
	"math"
)

var callID = 0

type Line struct {
	P1, P2 Position
}

func (c *Canvas) DrawLine(l Line, p Pixel) {
	l.P1.X = math.Floor(l.P1.X)
	l.P1.Y = math.Floor(l.P1.Y)
	l.P2.X = math.Floor(l.P2.X)
	l.P2.Y = math.Floor(l.P2.Y)

	dx := math.Abs(float64(l.P2.X - l.P1.X))
	dy := math.Abs(float64(l.P2.Y - l.P1.Y))

	// Fast cases
	if dx == 0 {
		c.drawVerticalLine(l, p)
		return
	} else if dy == 0 {
		c.drawHorizontalLine(l, p)
		return
	}

	var sx, sy float64
	if l.P1.X < l.P2.X {
		sx = 1
	} else {
		sx = -1
	}
	if l.P1.Y < l.P2.Y {
		sy = 1
	} else {
		sy = -1
	}

	err := dx - dy

	x := l.P1.X
	y := l.P1.Y
	for {
		c.SetPixel(int(x), int(y), p)
		if x == l.P2.X && y == l.P2.Y {
			break
		}
		e2 := 2 * err
		if e2 > -dy {
			err -= dy
			x += sx
		}
		if e2 < dx {
			err += dx
			y += sy
		}
	}
}

func (c *Canvas) drawHorizontalLine(l Line, p Pixel) {
	if l.P1.X > l.P2.X {
		l.P1, l.P2 = l.P2, l.P1
	}
	for x := l.P1.X; x <= l.P2.X; x++ {
		c.SetPixel(int(x), int(l.P1.Y), p)
	}
}

func (c *Canvas) drawVerticalLine(l Line, p Pixel) {
	if l.P1.Y > l.P2.Y {
		l.P1, l.P2 = l.P2, l.P1
	}
	for y := l.P1.Y; y <= l.P2.Y; y++ {
		c.SetPixel(int(l.P1.X), int(y), p)
	}
}
