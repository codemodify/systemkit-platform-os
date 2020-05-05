package os

const (
	OS_Gentoo         Name = "Gentoo"          // Gentoo (linux), Gentoo/FreeBSD (freebsd)
	OS_CalculateLinux Name = "Calculate Linux" //
	OS_ChromeOS       Name = "Chrome OS"       //
	OS_ContainerLinux Name = "Container Linux" //
	OS_Gentoox        Name = "Gentoox"         //
	OS_Knopperdisk    Name = "Knopperdisk"     //
	OS_Pentoo         Name = "Pentoo"          //
	OS_SabayonLinux   Name = "Sabayon Linux"   //
	OS_SystemRescueCD Name = "SystemRescueCD"  //
	OS_TinHatLinux    Name = "Tin Hat Linux"   //
	OS_Ututo          Name = "Ututo"           //
)

var osNameFamilyLinuxGentoo = []Name{
	OS_Gentoo,
	OS_CalculateLinux,
	OS_ChromeOS,
	OS_ContainerLinux,
	OS_Gentoox,
	OS_Knopperdisk,
	OS_Pentoo,
	OS_SabayonLinux,
	OS_SystemRescueCD,
	OS_TinHatLinux,
	OS_Ututo,
}
