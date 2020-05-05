// +build linux

package os

func getFamilyNameVersion_Debian_Ubuntu_Knoppix(lines []string) (Family, Name, string) {
	familyName := Family_Unknown
	osName := OS_Unknown
	version := ""

	// cat /etc/*-release

	// Raspbian
	// ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~
	// PRETTY_NAME = "Raspbian GNU/Linux 10 (buster)"
	// NAME = "Raspbian GNU/Linux"
	// VERSION_ID = "10"
	// VERSION = "10 (buster)"
	// VERSION_CODENAME = buster
	// ID = raspbian
	// ID_LIKE = debian
	// HOME_URL = "http://www.raspbian.org/"
	// SUPPORT_URL = "http://www.raspbian.org/RaspbianForums"
	// BUG_REPORT_URL = "http://www.raspbian.org/RaspbianBugs"

	// Ubuntu
	// ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~

	// detect family
	for _, line := range lines {
		if fetchWordAt(line, "=", 0) == "NAME" {
			fullOSNameAsString := fetchWordAt(line, "=", 1)

			// Debian
			if isInList, name := checkOsInOsGroup(fullOSNameAsString, osNameFamilyLinuxDebian); isInList {
				familyName = Family_Linux_Debian
				osName = name
			}

			// Ubuntu
			if isInList, name := checkOsInOsGroup(fullOSNameAsString, osNameFamilyLinuxUbuntu); isInList {
				familyName = Family_Linux_Ubuntu
				osName = name
			}

			// Knoppix
			if isInList, name := checkOsInOsGroup(fullOSNameAsString, osNameFamilyLinuxKnoppix); isInList {
				familyName = Family_Linux_Knoppix
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
		getFamilyNameVersion_Debian_Ubuntu_Knoppix,
	)
}
