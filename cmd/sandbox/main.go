package main

import (
	"os"
	"time"

	"github.com/Polo123456789/go-game/pkg/trender"
	"github.com/Polo123456789/go-game/pkg/trender/8-color-pixels"
)

const (
	Width  = 80
	Height = 24
)

func main() {
	os.Stdout.WriteString(trender.AnsiClearScreen)

	defaultPixel := pixels.Pixel{
		Background:   pixels.White,
		Foreground:   pixels.Black,
		Content:      ' ',
		GraphicsMode: trender.Reset,
	}

	canvas := trender.NewCanvas(
		Width,
		Height,
		&defaultPixel,
	)

	ballPixel := pixels.Pixel{
		Background:   pixels.Red,
		Foreground:   pixels.Black,
		Content:      ' ',
		GraphicsMode: trender.Reset,
	}

	ballX := float64(Width / 2)
	ballY := float64(1)
	ballVelocity := 0.0

	gravity := 0.2

	for {
		canvas.SetPixel(
			int(ballX),
			int(ballY),
			&ballPixel,
		)

		canvas.FullRender()

		canvas.SetPixel(
			int(ballX),
			int(ballY),
			&defaultPixel,
		)

		ballY += ballVelocity
		ballVelocity += gravity

		if ballY < 0 {
			ballY = 0
			ballVelocity = -ballVelocity
		}

		if ballY >= Height {
			ballY = Height - 1
			ballVelocity = -ballVelocity * 0.8
		}

		time.Sleep(36 * time.Millisecond)
	}
}
