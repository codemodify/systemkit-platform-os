// +build linux
// +build !android

package tests

import (
	"fmt"
	"strings"
	"syscall"
	"testing"
)

func TestCPUVariant(t *testing.T) {
	var uname syscall.Utsname
	if err := syscall.Uname(&uname); err != nil {
		fmt.Printf("Uname: %v", err)
	}
	fmt.Println("uname.Sysname   : " + arrayToString(uname.Sysname))
	fmt.Println("uname.Nodename  : " + arrayToString(uname.Nodename))
	fmt.Println("uname.Release   : " + arrayToString(uname.Release))
	fmt.Println("uname.Version   : " + arrayToString(uname.Version))
	fmt.Println("uname.Machine   : " + arrayToString(uname.Machine))
	fmt.Println("uname.Domainname: " + arrayToString(uname.Domainname))
}

func arrayToString(x [65]int8) string {
	var buf [65]byte
	for i, b := range x {
		buf[i] = byte(b)
	}
	str := string(buf[:])
	if i := strings.Index(str, "\x00"); i != -1 {
		str = str[:i]
	}
	return str
}
