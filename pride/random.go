package pride

import (
	"fmt"
	"math/rand"
	"time"
)

// PrintRandomPrideFlag prints a random pride flag in your terminal.
func PrintRandomPrideFlag() {
	rand.Seed(time.Now().UnixNano())

	flagsLen := len(Flags)
	randFlagIndex := rand.Intn(flagsLen)
	currFlagIndex := 0

	for pride, printPrideFlag := range Flags {
		if currFlagIndex == randFlagIndex {
			fmt.Println(pride)
			printPrideFlag()
			return
		}
		currFlagIndex++
	}
}
