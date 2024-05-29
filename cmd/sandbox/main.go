package main

import (
	"fmt"
	"os"
	"time"

	"github.com/Polo123456789/go-game/pkg/trender"
	"github.com/Polo123456789/go-game/pkg/trender/256-color-pixels"
)

const (
	Width  = 60
	Height = 24
)

func main() {
	os.Stdout.WriteString(trender.AnsiClearScreen)

	defaultPixel := pixels.NewPixel(
		pixels.DefaultColor,
		pixels.ColorID(4),
		' ',
	)
	canvas := trender.NewCanvas(
		Width,
		Height,
		defaultPixel,
	)

	ballPixel := pixels.NewPixel(
		pixels.DefaultColor,
		pixels.ColorID(1),
		' ',
	)

	ballX := float64(Width / 2)
	ballY := float64(1)
	ballVelocity := 0.0

	gravity := 0.2
	start := time.Now()
	time.Sleep(100 * time.Millisecond)

	timeSpentSum := int64(0)
	framesCount := 0
	maxFrames := 10000

	for {

		canvas.SetPixel(
			int(ballX),
			int(ballY),
			ballPixel,
		)

		framesCount++
		start = time.Now()
		canvas.Render()
		timeSpentSum += time.Since(start).Nanoseconds()

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
