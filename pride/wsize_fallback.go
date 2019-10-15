// +build !linux,!darwin

package pride

func windowSize() (width, height int) {
	return 80, 24
}
