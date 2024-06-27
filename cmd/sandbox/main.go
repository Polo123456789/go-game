package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"runtime/pprof"

	"github.com/Polo123456789/go-game/pkg/trender"
	"github.com/Polo123456789/go-game/pkg/trender/8-color-pixels"
)

func run(ctx context.Context) error {
	trender.ClearScreen()
	size, err := trender.GetTermSize()
	if err != nil {
		return err
	}

	size.Height -= 1

	background := pixels.NewPixel(
		pixels.White,
		pixels.White,
		pixels.Reset,
		' ',
	)

	lineForeground := pixels.NewPixel(
		pixels.Black,
		pixels.Black,
		pixels.Reset,
		' ',
	)

	fillRectForeground := pixels.NewPixel(
		pixels.Yellow,
		pixels.Yellow,
		pixels.Reset,
		' ',
	)

	rectForeground := pixels.NewPixel(
		pixels.Red,
		pixels.Red,
		pixels.Reset,
		' ',
	)

	elipseForeground := pixels.NewPixel(
		pixels.Green,
		pixels.Green,
		pixels.Reset,
		' ',
	)
	// TODO: Implement elipse drawing
	_ = elipseForeground

	canvas := trender.NewCanvas(
		size.Width,
		size.Height,
		background,
	)

	lines := []trender.Line{
		{
			P1: trender.Position{X: 0, Y: 0},
			P2: trender.Position{X: float64(size.Width - 1), Y: float64(size.Height - 1)},
		},
		{
			P1: trender.Position{X: float64(size.Width - 1), Y: 0},
			P2: trender.Position{X: 0, Y: float64(size.Height - 1)},
		},
		{
			P2: trender.Position{X: float64(size.Width / 2), Y: 0},
			P1: trender.Position{X: float64(size.Width / 2), Y: float64(size.Height - 1)},
		},
		{
			P2: trender.Position{X: 0, Y: float64(size.Height / 2)},
			P1: trender.Position{X: float64(size.Width - 1), Y: float64(size.Height / 2)},
		},
	}
	for _, line := range lines {
		canvas.DrawLine(line, lineForeground)
	}

	rect := trender.Rect{
		Center: trender.Position{
			X: float64(size.Width / 2),
			Y: float64(size.Height / 2),
		},
		Width:  20,
		Height: 10,
	}
	canvas.DrawRectFill(rect, fillRectForeground)

	noEmptyRects := 10
	for i := 0; i < noEmptyRects; i++ {
		rect := trender.Rect{
			Center: trender.Position{
				X: float64(size.Width / 2),
				Y: float64(size.Height / 2),
			},
			Width:  20 + i*5,
			Height: 10 + i*5,
		}
		canvas.DrawRect(rect, rectForeground)
	}

	canvas.RenderFull()

	return nil
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	cpuFile, err := os.Create("cpu.prof")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	pprof.StartCPUProfile(cpuFile)
	defer pprof.StopCPUProfile()

	memFile, err := os.Create("mem.prof")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	defer pprof.WriteHeapProfile(memFile)

	if err := run(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
