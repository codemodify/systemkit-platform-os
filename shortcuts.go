package os

import (
	"fmt"
	"strings"
)

// IsDarwin -
func IsDarwin(name OSName) bool {
	return (name == OS_Darwin || name == OS_MacOSX)
}

// IsBSD -
func IsBSD(name OSName) bool {
	if strings.Contains(strings.ToLower(fmt.Sprint("", name)), "bsd") {
		return true
	}

	return (name == OS_Dragonfly || name == OS_Bitrig)
}

// IsWindows -
func IsWindows(name OSName) bool {
	return (name == OS_Windows ||
		name == OS_CYGWIN ||
		name == OS_MSYS ||
		name == OS_MINGW)
}
