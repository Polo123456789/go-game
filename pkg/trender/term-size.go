package trender

import (
	"fmt"
	"os"
	"os/exec"
)

type TermSize struct {
	Width  int
	Height int
}

func GetTermSize() (TermSize, error) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		return TermSize{}, err
	}
	var ts TermSize
	n, err := fmt.Sscanf(string(out), "%d %d", &ts.Height, &ts.Width)
	if err != nil {
		return TermSize{}, err
	}
	if n != 2 {
		return TermSize{}, fmt.Errorf("expected 2 values, got %d", n)
	}
	return ts, nil
}
