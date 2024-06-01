package main

import (
	"github.com/Polo123456789/go-game/pkg/trender/256-color-pixels"
)

const (
	MIN_ASTEROID_RADIUS = 2
	MAX_ASTEROID_RADIUS = 5
)

// Should be constant, but Go doesn't allow it
var ASTEROID_COLORS = [...]pixels.ColorID{
	pixels.ColorID(245),
	pixels.ColorID(240),
	pixels.ColorID(238),
}

type Asteroid struct {
	radius       float64
	rotation     float64
	x, y, vx, vy float64
}
