package pride

import (
	"fmt"
)

// PrintNonbinaryPrideFlag prints a nonbinary pride flag in your terminal.
func PrintNonbinaryPrideFlag() {
	_, height := windowSize()
	segmentHeight := height / 4
	for segment := 1; segment <= 4; segment++ {
		line := "\n"
		for i := 0; i < segmentHeight-1; i++ {
			line += "\n"
		}
		if segment == 4 {
			line += "\n"
		}

		var colorCode string
		if segment == 1 {
			colorCode = "011"
		}
		if segment == 2 {
			colorCode = "015"
		}
		if segment == 3 {
			colorCode = "054"
		}
		if segment == 4 {
			colorCode = "000"
		}
		fmt.Printf("\x1B[48;5;%sm%s\x1B[0m", colorCode, line)
	}
}
