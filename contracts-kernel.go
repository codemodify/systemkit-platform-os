package os

type Kernel string

// https://en.wikipedia.org/wiki/Comparison_of_operating_system_kernels
const (
	Kernel_Amiga        Kernel = "Amiga"        //
	Kernel_DragonFlyBSD Kernel = "DragonFlyBSD" //
	Kernel_FreeBSD      Kernel = "FreeBSD"      //
	Kernel_GNUHurd      Kernel = "GNUHurd"      //
	Kernel_GNUMatch     Kernel = "GNUMatch"     //
	Kernel_Inferno      Kernel = "Inferno"      //
	Kernel_Illumos      Kernel = "Illumos"      //
	Kernel_L4           Kernel = "L4"           //
	Kernel_Linux        Kernel = "Linux"        //
	Kernel_Match        Kernel = "Match"        //
	Kernel_MINIX        Kernel = "MINIX"        //
	Kernel_MkLinux      Kernel = "MkLinux"      //
	Kernel_NetBSD       Kernel = "NetBSD"       //
	Kernel_NetWare      Kernel = "NetWare"      //
	Kernel_OpenBSD      Kernel = "OpenBSD"      //
	Kernel_OS2          Kernel = "OS2"          //
	Kernel_Plan9        Kernel = "Plan9"        //
	Kernel_ReactOS      Kernel = "ReactOS"      //
	Kernel_Rockbox      Kernel = "Rockbox"      //
	Kernel_SunOS        Kernel = "SunOS"        //
	Kernel_Solaris      Kernel = "Solaris"      //
	Kernel_Trix         Kernel = "Trix"         //
	Kernel_WindowsNT    Kernel = "WindowsNT"    //
	Kernel_XNU          Kernel = "XNU"          // Darwin
	Kernel_SPARTAN      Kernel = "SPARTAN"      //
	Kernel_Unknown      Kernel = "Unknown"      //
)
