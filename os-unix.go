// +build !windows

package os

import (
	"strings"
)

// Info -
func Info() OS {
	// collect all
	osInfos := []OS{
		tryMethod_uname(),
		tryMethod_hostnamectl(),
		tryMethod_hostnamectl(),
	}

	// merge
	result := OS{}

	for _, osInfo := range osInfos {
		if osInfo.Name != OS_Uknown && len(osInfo.Name) > 0 && len(result.Name) == 0 {
			result.Name = osInfo.Name
		}

		if len(osInfo.Version) > 0 && len(result.Version) == 0 {
			result.Version = osInfo.Version
		}

		if len(osInfo.Distribution) > 0 && len(result.Distribution) == 0 {
			result.Distribution = osInfo.Distribution
		}

		if len(osInfo.Features) > 0 {
			result.Features = append(result.Features, osInfo.Features...)
		}
	}

	if len(result.Name) == 0 {
		result.Name = OS_Uknown
	}

	return result
}

func intArrayToString(x [65]int8) string {
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

func tryMethod_uname() OS {
	// var uname syscall.Utsname
	// if err := syscall.Uname(&uname); err != nil {
	// 	return OS{}, err
	// }

	// return OS{
	// 	Name:         getOSFromString(intArrayToString(uname.Sysname)),
	// 	Version:      intArrayToString(uname.Release),
	// 	Distribution: "",
	// 	Features:     []string{},
	// }, nil

	output1, err1 := runBinaryFetchOutput("uname", []string{"-s"})
	if err1 != nil {
		return OS{}
	}

	output2, err2 := runBinaryFetchOutput("uname", []string{"-r"})
	if err2 != nil {
		return OS{}
	}

	return OS{
		Name:         getOSFromString(fetchFirstInList(strings.Join(output1, ""), "")),
		Version:      fetchFirstInList(strings.Join(output2, ""), ""),
		Distribution: "",
		Features:     []string{},
	}
}

func tryMethod_hostnamectl() OS {
	//
	// hostnamectl - part of SystemD, will for sure fail on distros not using it
	//

	output, err := runBinaryFetchOutput("hostnamectl", []string{"status"})
	if err != nil {
		return OS{}
	}

	return OS{
		Name:         getOSFromString(fetchFirstInList(fetchFromLineToRight(output, "Kernel", ":"), " ")), // ex: Linux
		Version:      fetchSecondInList(fetchFromLineToRight(output, "Kernel", ":"), " "),                 // ex: 5.5.4-arch1-1
		Distribution: fetchFromLineToRight(output, "Operating System", ":"),                               // ex: Arch Linux
		Features:     []string{},
	}
}

func tryMethod_etc_os_release() OS {
	//
	// should work on most modern linuxes
	//

	output, err := runBinaryFetchOutput("cat", []string{"/etc/*-release"})
	if err != nil {
		return OS{}
	}

	return OS{
		Name:         getOSFromString(cleanUp(fetchFromLineToRight(output, "NAME", "="), []string{"\"", "GNU", "/"})),
		Version:      "",
		Distribution: cleanUp(fetchFromLineToRight(output, "PRETTY_NAME", "="), []string{"\"", "GNU", "/"}),
		Features:     []string{},
	}

	// #
	// # RPI
	// #
	// PRETTY_NAME="Raspbian GNU/Linux 10 (buster)"
	// NAME="Raspbian GNU/Linux"
	// VERSION_ID="10"
	// VERSION="10 (buster)"
	// VERSION_CODENAME=buster
	// ID=raspbian
	// ID_LIKE=debian
	// HOME_URL="http://www.raspbian.org/"
	// SUPPORT_URL="http://www.raspbian.org/RaspbianForums"
	// BUG_REPORT_URL="http://www.raspbian.org/RaspbianBugs"
	// #
	// # Arch Linux
	// #
	// NAME="Arch Linux"
	// PRETTY_NAME="Arch Linux"
	// ID=arch
	// BUILD_ID=rolling
	// ANSI_COLOR="0;36"
	// HOME_URL="https://www.archlinux.org/"
	// DOCUMENTATION_URL="https://wiki.archlinux.org/"
	// SUPPORT_URL="https://bbs.archlinux.org/"
	// BUG_REPORT_URL="https://bugs.archlinux.org/"
	// LOGO=archlinux
	// #
	// # Ubuntu Linux v7.10 server
	// #
	// DISTRIB_ID=Ubuntu
	// DISTRIB_RELEASE=7.10
	// DISTRIB_CODENAME=gutsy
	// DISTRIB_DESCRIPTION="Ubuntu 7.10"

	// There's at least
	// /etc/redhat-release
	// /etc/SuSE-release
	// /etc/debian_version
	// /etc/arch-release
	// /etc/gentoo-release
	// /etc/slackware-version
	// /etc/frugalware-release
	// /etc/altlinux-release
	// /etc/mandriva-release
	// /etc/meego-release
	// /etc/angstrom-version
	// /etc/mageia-release
}
