package pride

import (
	"fmt"
)

// PrintPansexualPrideFlag prints a pansexual pride flag in your terminal.
func PrintPansexualPrideFlag() {
	_, height := windowSize()
	segmentHeight := height / 3
	for segment := 1; segment <= 3; segment++ {
		line := "\n"
		for i := 0; i < segmentHeight-1; i++ {
			line += "\n"
		}
		if segment == 3 {
			line += "\n"
		}

		var colorCode string
		if segment == 1 {
			colorCode = "197"
		}
		if segment == 2 {
			colorCode = "220"
		}
		if segment == 3 {
			colorCode = "038"
		}
		fmt.Printf("\x1B[48;5;%sm%s\x1B[0m", colorCode, line)
	}
}
