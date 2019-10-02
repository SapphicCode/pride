package pride

import (
  "os"
  "golang.org/x/crypto/ssh/terminal"
)

func windowSize() (height, width int) {
  width, height, err := terminal.GetSize(int(os.Stdout.Fd()))
  if err != nil {
    return 24, 80
  }
  return height, width
}
