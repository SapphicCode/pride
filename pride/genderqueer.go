package pride

import (
	"fmt"
)

// PrintGenderqueerPrideFlag prints a genderqueer pride flag in your terminal.
func PrintGenderqueerPrideFlag() {
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

		var colorCode int
		switch segment {
		case 1:
			colorCode = 105
		case 2:
			colorCode = 107
		case 3:
			colorCode = 42
		}
		fmt.Printf("\x1B[%dm%s\x1B[0m", colorCode, line)
	}
}
