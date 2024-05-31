package main

import (
	"fmt"
	"os"
	"runtime/pprof"
	"time"

	"github.com/Polo123456789/go-game/pkg/trender"
	"github.com/Polo123456789/go-game/pkg/trender/rgb-true-color-pixels"
)

const (
	Width  = 180
	Height = 60
)

func run() error {
	cpuProfileFile, err := os.Create("cpu.prof")
	if err != nil {
		return err
	}
	pprof.StartCPUProfile(cpuProfileFile)
	defer pprof.StopCPUProfile()

	memProfileFile, err := os.Create("mem.prof")
	if err != nil {
		return err
	}
	defer pprof.WriteHeapProfile(memProfileFile)

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
		canvas.RenderFull()
		timeSpentSum += time.Since(start).Milliseconds()

		if framesCount >= maxFrames {
			os.Stdout.WriteString(trender.SetCursorPosition(0, 27))
			fmt.Printf("Frames Count: %v\n", framesCount)
			fmt.Printf("Average TimeSpent on Render: %v", timeSpentSum/int64(maxFrames))
			return nil
		}

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

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
