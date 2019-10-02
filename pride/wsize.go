package pride

import (
	"golang.org/x/crypto/ssh/terminal"
	"os"
)

func windowSize() (height, width int) {
	width, height, err := terminal.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return 24, 80
	}
	return height, width
}
