// +build linux darwin

package pride

import (
	"golang.org/x/crypto/ssh/terminal"
	"os"
)

func windowSize() (width, height int) {
	width, height, err := terminal.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return 80, 24
	}
	return
}
