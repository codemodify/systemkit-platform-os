// +build linux

package os

import "strings"

func info() OS {
	result := newOS()

	// 1 - Family + Name + Version
	result.Family, result.Name, result.Version = getFamilyNameVersion()

	// 2 - Kernel
	result.Kernel.Name = Kernel_Linux
	result.Kernel.Version = getKernelVersion()
	result.Kernel.BuildTime = getKernelBuildTime()
	result.Kernel.ExecFormat = ExecFormat_ELF

	// 3 - Features

	return result
}

var ignoreWords = []string{
	"linux",
}

func checkOsInOsGroup(osVal string, osGroup []Name) (bool, Name) {
	osVal = removeWords(osVal, []string{"\""})
	osVal = strings.ToLower(osVal)

	for _, osName := range osGroup {
		familyNameAsString := strings.ToLower(string(osName))
		familyNameAsString = removeWords(familyNameAsString, ignoreWords)

		if familyNameAsString == fetchWordAt(osVal, " ", 0) {
			return true, osName
		}
	}

	return false, OS_Unknown
}

type getFamilyNameVersionDelegate func(lines []string) (Family, Name, string)

var familyNameVersionDelegates = []getFamilyNameVersionDelegate{}

func getFamilyNameVersion() (Family, Name, string) {
	familyName := Family_Unknown
	osName := OS_Unknown
	version := ""

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
	// # Ubuntu Linux v7.10 server
	// #
	// DISTRIB_ID=Ubuntu
	// DISTRIB_RELEASE=7.10
	// DISTRIB_CODENAME=gutsy
	// DISTRIB_DESCRIPTION="Ubuntu 7.10"

	// There's at least
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
	output := runBinaryFetchOutput("sh", []string{"-c", "cat /etc/*-release"})
	for _, delegate := range familyNameVersionDelegates {
		familyName, osName, version = delegate(output)
		if familyName != Family_Unknown {
			break
		}
	}

	return familyName, osName, version
}

func getKernelVersion() string {
	output := runBinaryFetchOutput("uname", []string{"--kernel-release"})
	return fetchFirstInList(output)
}

func getKernelBuildTime() string {
	output := runBinaryFetchOutput("uname", []string{"--kernel-version"})
	return fetchFirstInList(output)
}
