package os

import (
	"fmt"
	"strings"
)

// IsDarwin -
func IsDarwin(name Name) bool {
	return (name == OS_Darwin || name == OS_MacOSX)
}

// IsBSD -
func IsBSD(name Name) bool {
	if strings.Contains(strings.ToLower(fmt.Sprint("", name)), "bsd") {
		return true
	}

	return (name == OS_DragonflyBSD || name == OS_Bitrig)
}

// IsWindows -
func IsWindows(name Name) bool {
	return (name == OS_Windows ||
		name == OS_CYGWIN ||
		name == OS_MSYS ||
		name == OS_MINGW)
}
