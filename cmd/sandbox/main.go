package main

import (
	"fmt"
	"os"
	"time"

	"github.com/Polo123456789/go-game/pkg/trender"
	"github.com/Polo123456789/go-game/pkg/trender/rgb-true-color-pixels"
)

const (
	Width  = 60
	Height = 24
)

func main() {
	os.Stdout.WriteString(trender.AnsiClearScreen)

	defaultPixel := pixels.NewPixel(
		pixels.RGB{
			R: 0,
			G: 0,
			B: 0,
		},
		pixels.RGB{
			R: 0xff,
			G: 0xff,
			B: 0xff,
		},
		' ',
	)
	canvas := trender.NewCanvas(
		Width,
		Height,
		defaultPixel,
	)

	ballPixel := pixels.NewPixel(
		pixels.RGB{
			R: 0,
			G: 0,
			B: 0,
		},
		pixels.RGB{
			R: 41,
			G: 205,
			B: 217,
		},
		' ',
	)

	ballX := float64(Width / 2)
	ballY := float64(1)
	ballVelocity := 0.0

	gravity := 0.2

	timeSpentSum := int64(0)
	framesCount := 0
	maxFrames := 1000

	for {

		canvas.SetPixel(
			int(ballX),
			int(ballY),
			ballPixel,
		)

		framesCount++
		start := time.Now()
		canvas.FullRender()
		timeSpentSum += time.Since(start).Milliseconds()

		if framesCount >= maxFrames {
			os.Stdout.WriteString(trender.SetCursorPosition(0, 27))
			fmt.Printf("Frames Count: %v\n", framesCount)
			fmt.Printf("Average TimeSpent on Render: %v", timeSpentSum/int64(maxFrames))
			return
		}

		start = time.Now()

		canvas.SetPixel(
			int(ballX),
			int(ballY),
			defaultPixel,
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
	}
}
