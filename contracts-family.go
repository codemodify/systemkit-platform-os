package os

type Family string

const (
	Family_AIX             Family = "AIX"          //
	Family_Darwin          Family = "Darwin"       //
	Family_DragonflyBSD    Family = "DragonflyBSD" //
	Family_FreeBSD         Family = "FreeBSD"      //
	Family_HURD            Family = "HURD"         //
	Family_Illumos         Family = "Illumos"      //
	Family_JS              Family = "JS"           //
	Family_Linux           Family = "Linux"        //
	Family_Linux_Android   Family = "Android"      //
	Family_Linux_ArchLinux Family = "Arch Linux"   //

	Family_Linux_Debian  Family = "Debian"  //
	Family_Linux_Knoppix Family = "Knoppix" //
	Family_Linux_Ubuntu  Family = "Ubuntu"  //

	Family_Linux_Gentoo Family = "Gentoo" //

	Family_Linux_RHEL   Family = "Red Hat Enterprise Linux" //
	Family_Linux_CentOS Family = "CentOS"                   //
	Family_Linux_Fedora Family = "Fedora"                   //

	Family_Linux_Slackware Family = "Slackware" //
	Family_Linux_Slax      Family = "Slax"      //

	Family_Linux_Suse Family = "Suse" //

	Family_NACL    Family = "NACL"    //
	Family_NetBSD  Family = "NetBSD"  //
	Family_OpenBSD Family = "OpenBSD" //
	Family_Plan9   Family = "Plan9"   //
	Family_Solaris Family = "Solaris" //
	Family_Windows Family = "Windows" //
	Family_ZOS     Family = "ZOS"     //

	Family_Unknown Family = "Unknown" //
)
