package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/Polo123456789/go-game/pkg/trender"
	"github.com/Polo123456789/go-game/pkg/trender/256-color-pixels"
)

const BACKGROUND_COLOR = pixels.ColorID(232)
const STARS_COLOR = pixels.ColorID(255)

func run(ctx context.Context) error {
	size, err := trender.GetTermSize()
	if err != nil {
		return err
	}

	canvas := trender.NewCanvas(
		size.Width,
		size.Height,
		pixels.NewPixel(
			STARS_COLOR,
			BACKGROUND_COLOR,
			' ',
		),
	)

	// TODO: Add geometric shapes to trender, then come back to make the game

	canvas.RenderFull()

	return ctx.Err()
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, os.Interrupt)
	signal.Notify(sigchan, os.Kill)
	go func() {
		<-sigchan
		cancel()
	}()

	if err := run(ctx); err != nil {
		log.Printf("Error: %v", err)
		os.Exit(1)
	}
}
