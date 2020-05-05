// +build linux

package os

import "strings"

func getFamilyNameVersion_RHEL_CentOS_Fedora(lines []string) (Family, Name, string) {
	familyName := Family_Unknown
	osName := OS_Unknown
	version := ""

	// cat /etc/*-release

	// CentOS
	// ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~
	// CentOS Linux release 7.6.1810 (Core)
	// NAME="CentOS Linux"
	// VERSION="7 (Core)"
	// ID="centos"
	// ID_LIKE="rhel fedora"
	// VERSION_ID="7"
	// PRETTY_NAME="CentOS Linux 7 (Core)"
	// ANSI_COLOR="0;31"
	// CPE_NAME="cpe:/o:centos:centos:7"
	// HOME_URL="https://www.centos.org/"
	// BUG_REPORT_URL="https://bugs.centos.org/"
	// CENTOS_MANTISBT_PROJECT="CentOS-7"
	// CENTOS_MANTISBT_PROJECT_VERSION="7"
	// REDHAT_SUPPORT_PRODUCT="centos"
	// REDHAT_SUPPORT_PRODUCT_VERSION="7"
	// CentOS Linux release 7.6.1810 (Core)
	// CentOS Linux release 7.6.1810 (Core)

	// detect family
	for _, line := range lines {
		if fetchWordAt(line, "=", 0) == "NAME" {
			fullOSNameAsString := fetchWordAt(line, "=", 1)

			// RHEL
			if isInList, name := checkOsInOsGroup(fullOSNameAsString, osNameFamilyLinuxRHEL); isInList {
				familyName = Family_Linux_RHEL
				osName = name
			}

			// CentOS
			if isInList, name := checkOsInOsGroup(fullOSNameAsString, osNameFamilyLinuxCentos); isInList {
				familyName = Family_Linux_CentOS
				osName = name
			}

			// Fedodra
			if isInList, name := checkOsInOsGroup(fullOSNameAsString, osNameFamilyLinuxFedora); isInList {
				familyName = Family_Linux_Fedora
				osName = name
			}

		} else if fetchWordAt(line, "=", 0) == "VERSION" {
			version = fetchWordAt(line, "=", 1)
			version = removeWords(version, []string{"\""})

		} else if !strings.Contains(line, "=") &&
			strings.Contains(line, string(osName)) &&
			strings.Contains(line, "release") {
			version = removeWords(line, []string{
				string(osName),
				"release",
			})
		}
	}

	return familyName, osName, version
}

func init() {
	familyNameVersionDelegates = append(
		familyNameVersionDelegates,
		getFamilyNameVersion_RHEL_CentOS_Fedora,
	)
}
