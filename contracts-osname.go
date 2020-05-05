package os

type Name string

const OS_Unknown Name = "Unknown"

var supportedOses = []Name{}

func init() {
	supportedOses = append(supportedOses, osNameFamilyAIX...)
	supportedOses = append(supportedOses, osNameFamilyDarwin...)
	supportedOses = append(supportedOses, osNameFamilyDragonflyBSD...)
	supportedOses = append(supportedOses, osNameFamilyFreeBSD...)
	supportedOses = append(supportedOses, osNameFamilyHURD...)
	supportedOses = append(supportedOses, osNameFamilyIllumos...)
	supportedOses = append(supportedOses, osNameFamilyJS...)
	supportedOses = append(supportedOses, osNameFamilyLinuxAndroid...)
	supportedOses = append(supportedOses, osNameFamilyLinuxArch...)
	supportedOses = append(supportedOses, osNameFamilyLinuxCentos...)
	supportedOses = append(supportedOses, osNameFamilyLinuxDebian...)
	supportedOses = append(supportedOses, osNameFamilyLinuxFedora...)
	supportedOses = append(supportedOses, osNameFamilyLinuxGentoo...)
	supportedOses = append(supportedOses, osNameFamilyLinuxKnoppix...)
	supportedOses = append(supportedOses, osNameFamilyLinuxSlackware...)
	supportedOses = append(supportedOses, osNameFamilyLinuxSlax...)
	supportedOses = append(supportedOses, osNameFamilyLinuxSuse...)
	supportedOses = append(supportedOses, osNameFamilyLinuxUbuntu...)
	supportedOses = append(supportedOses, osNameFamilyLinux...)
	supportedOses = append(supportedOses, osNameFamilyNACL...)
	supportedOses = append(supportedOses, osNameFamilyNetBSD...)
	supportedOses = append(supportedOses, osNameFamilyOpenBSD...)
	supportedOses = append(supportedOses, osNameFamilyOther...)
	supportedOses = append(supportedOses, osNameFamilyPlan9...)
	supportedOses = append(supportedOses, osNameFamilySolaris...)
	supportedOses = append(supportedOses, osNameFamilyWindows...)
	supportedOses = append(supportedOses, osNameFamilyZOS...)
}
