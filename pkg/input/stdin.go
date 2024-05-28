package input

import (
	"os"
	"os/exec"
)

func UnbufferStdin() {
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
}

func RestoreStdin() {
	exec.Command("stty", "-F", "/dev/tty", "echo").Run()
}

func Get() rune {
	b := [1]byte{}
	os.Stdin.Read(b[:])
	return rune(b[0])
}
