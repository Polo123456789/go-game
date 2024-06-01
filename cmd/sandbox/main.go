package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"runtime/pprof"
	"time"

	"github.com/Polo123456789/go-game/pkg/trender"
	"github.com/Polo123456789/go-game/pkg/trender/8-color-pixels"
)

type Particle struct {
	x, y, vx, vy float64
}

func NewRandomParticle(width, height int) Particle {
	return Particle{
		x:  float64(rand.Intn(width)),
		y:  float64(rand.Intn(height)),
		vx: rand.Float64()*2 - 1,
		vy: rand.Float64()*2 - 1,
	}
}

type Simulation struct {
	Canvas     *trender.Canvas
	Particles  []Particle
	Elasticity float64

	background trender.Pixel
	particle   trender.Pixel
}

func NewSimulation(
	width, height int,
	elasticity float64,
	particleCount int,
	backgroundPixel trender.Pixel,
	particlePixel trender.Pixel,
) *Simulation {
	s := &Simulation{
		Canvas:     trender.NewCanvas(width, height, backgroundPixel),
		Elasticity: elasticity,
		background: backgroundPixel,
		particle:   particlePixel,
	}
	for i := 0; i < particleCount; i++ {
		s.Particles = append(s.Particles, NewRandomParticle(width, height))
	}
	return s
}

func (s *Simulation) DrawParticles(pixel trender.Pixel) {
	for _, p := range s.Particles {
		s.Canvas.SetPixel(int(p.x), int(p.y), pixel)
	}
}

func (s *Simulation) UpdateParticles() {
	for i := range s.Particles {
		p := &s.Particles[i]
		p.x += p.vx
		p.y += p.vy
		if p.x < 0 {
			p.x = 0
			p.vx = -p.vx * s.Elasticity
		}
		if p.x >= float64(s.Canvas.Width()) {
			p.x = float64(s.Canvas.Width()) - 1
			p.vx = -p.vx * s.Elasticity
		}
		if p.y < 0 {
			p.y = 0
			p.vy = -p.vy * s.Elasticity
		}
		if p.y >= float64(s.Canvas.Height()) {
			p.y = float64(s.Canvas.Height()) - 1
			p.vy = -p.vy * s.Elasticity
		}
	}
}

func runSimulation(ctx context.Context, args []string) error {
	flags := flag.NewFlagSet("game", flag.ExitOnError)
	particleCount := flags.Int("particles", 1000, "Number of particles")
	elasticity := flags.Float64("elasticity", 0.9, "Elasticity of the particles")
	flags.Parse(args[1:])

	size, err := trender.GetTermSize()
	if err != nil {
		return err
	}

	s := NewSimulation(
		size.Width,
		size.Height,
		*elasticity,
		*particleCount,
		pixels.NewPixel(
			pixels.Black,
			pixels.White,
			pixels.Reset,
			' ',
		),
		pixels.NewPixel(
			pixels.White,
			pixels.Black,
			pixels.Reset,
			' ',
		),
	)

	start := time.Now()
	simStart := time.Now()
	var renderTimeSum time.Duration
	var frameCount int

	for ctx.Err() == nil {
		s.DrawParticles(s.background)
		s.UpdateParticles()
		s.DrawParticles(s.particle)
		spent := time.Since(start)
		frameCount++
		renderTimeSum += spent
		s.Canvas.RenderChanged()
		start = time.Now()
	}

	fmt.Printf("Frames rendered: %d\n", frameCount)
	fmt.Printf("Average render time: %v\n", renderTimeSum/time.Duration(frameCount))
	fmt.Printf("Simulation time: %v\n", time.Since(simStart))

	return nil
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, os.Interrupt)
	go func() {
		<-sigchan
		cancel()
	}()

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

	if err := runSimulation(ctx, os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
