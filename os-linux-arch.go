// +build linux

package os

func getFamilyNameVersion_Arch(lines []string) (Family, Name, string) {
	familyName := Family_Unknown
	osName := OS_Unknown
	version := ""

	// cat /etc/*-release

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

	// detect family
	for _, line := range lines {
		if fetchWordAt(line, "=", 0) == "NAME" {
			fullOSNameAsString := fetchWordAt(line, "=", 1)
			if isInList, name := checkOsInOsGroup(fullOSNameAsString, osNameFamilyLinuxArch); isInList {
				familyName = Family_Linux_ArchLinux
				osName = name
			}
		} else if fetchWordAt(line, "=", 0) == "VERSION" {
			version = fetchWordAt(line, "=", 1)
			version = removeWords(version, []string{"\""})
		}
	}

	return familyName, osName, version
}

func init() {
	familyNameVersionDelegates = append(
		familyNameVersionDelegates,
		getFamilyNameVersion_Arch,
	)
}
