package trender

import (
	"math"
)

type Line struct {
	P1, P2 Position
}

func (c *Canvas) DrawLine(l Line, p Pixel) {
	dx := math.Abs(float64(l.P2.X - l.P1.X))
	dy := math.Abs(float64(l.P2.Y - l.P1.Y))

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
