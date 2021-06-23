package prompt

import (
	"os"
	"os/exec"
)

type inputManager struct {
	C chan []byte
}

func newInputManager() *inputManager {
	// disable input buffering
	exec.Command("stty", "-f", "/dev/tty", "cbreak", "min", "1").Run()
	// do not display entered characters on the screen
	exec.Command("stty", "-f", "/dev/tty", "-echo").Run()
	c := make(chan []byte, 1)
	go func() {
		for {
			b := make([]byte, 3)
			os.Stdin.Read(b)
			c <- b
		}
	}()
	return &inputManager{
		C: c,
	}
}

func isEscapeSequence(b []byte) bool {
	if len(b) < 2 {
		return false
	}
	return b[0] == 27 && b[1] == 91
}
func isUpArrow(b []byte) bool {
	return isEscapeSequence(b[:2]) && b[2] == 65
}
func isDownArrow(b []byte) bool {
	return isEscapeSequence(b[:2]) && b[2] == 66
}
func isSpace(b []byte) bool {
	return b[0] == 32
}
func isBackspace(b []byte) bool {
	return b[0] == 8 || b[0] == 127
}
func isEnter(b []byte) bool {
	return b[0] == 10
}
func reduceInput(b []byte) []byte {
	r := make([]byte, 0)
	for _, _b := range b {
		if _b != 0 {
			r = append(r, _b)
		}
	}
	return r
}
