package trender

import ()

type Rect struct {
	Center Position
	Width  int
	Height int
}

func (c *Canvas) DrawRect(r Rect, p Pixel) {
	topLeft := Position{r.Center.X - float64(r.Width/2), r.Center.Y - float64(r.Height/2)}
	topRight := Position{r.Center.X + float64(r.Width/2), r.Center.Y - float64(r.Height/2)}
	bottomLeft := Position{r.Center.X - float64(r.Width/2), r.Center.Y + float64(r.Height/2)}
	bottomRight := Position{r.Center.X + float64(r.Width/2), r.Center.Y + float64(r.Height/2)}

	c.DrawLine(Line{topLeft, topRight}, p)
	c.DrawLine(Line{topRight, bottomRight}, p)
	c.DrawLine(Line{bottomRight, bottomLeft}, p)
	c.DrawLine(Line{bottomLeft, topLeft}, p)
}

func (c *Canvas) DrawRectFill(r Rect, p Pixel) {
	topLeft := Position{r.Center.X - float64(r.Width/2), r.Center.Y - float64(r.Height/2)}

	for x := 0; x < r.Width; x++ {
		top := Position{topLeft.X + float64(x), topLeft.Y}
		bottom := Position{topLeft.X + float64(x), topLeft.Y + float64(r.Height-1)}
		c.DrawLine(Line{top, bottom}, p)
	}
}
