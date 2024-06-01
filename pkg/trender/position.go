package trender

type Position struct {
	X, Y float64
}

func (p Position) Add(p2 Position) Position {
	return Position{p.X + p2.X, p.Y + p2.Y}
}
