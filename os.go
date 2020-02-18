package os

// OSName -
type OSName string

// OS -
type OS struct {
	Name         OSName   `json:"name,omitempty"`        // ex: Linux
	Version      string   `json:"version,omitempty"`     // ex: `10.0.14393.1066` on Windows
	Distribution string   `json:"dstribution,omitempty"` // ex: Arch
	Features     []string `json:"features,omitempty"`    // ex: `win32k` on Windows
}

const (
	// Supported out of the box by GOOS
	// https://github.com/golang/go/blob/master/src/go/build/syslist.go
	OS_AIX       OSName = "AIX"       // https://en.wikipedia.org/wiki/IBM_AIX
	OS_Android          = "Android"   // https://en.wikipedia.org/wiki/Android_app
	OS_Darwin           = "Darwin"    // https://en.wikipedia.org/wiki/Darwin_(operating_system)
	OS_Dragonfly        = "Dragonfly" // https://www.dragonflybsd.org/
	OS_FreeBSD          = "FreeBSD"   // https://www.freebsd.org/
	OS_HURD             = "HURD"      // https://www.gnu.org/software/hurd
	OS_Illumos          = "Illumos"   // https://www.illumos.org
	OS_JS               = "JS"        // https://github.com/os-js
	OS_Linux            = "Linux"     // https://www.kernel.org
	OS_NACL             = "NACL"      // https://en.wikipedia.org/wiki/Google_Native_Client
	OS_NetBSD           = "NetBSD"    // https://netbsd.org
	OS_OpenBSD          = "OpenBSD"   // https://www.openbsd.org
	OS_Plan9            = "Plan9"     // https://9p.io/plan9
	OS_Solaris          = "Solaris"   // https://www.oracle.com/solaris/solaris11
	OS_Windows          = "Windows"   // https://www.microsoft.com/en-us/windows
	OS_ZOS              = "ZOS"       // https://en.wikipedia.org/wiki/Z/OS, https://www.ibm.com/it-infrastructure/z/zos

	// Supported in addition to that
	OS_MacOSX   = "MacOSX"
	OS_iOS      = "iOS"
	OS_OS2      = "OS2"
	OS_Haiku    = "Haiku"
	OS_BeOS     = "BeOS"
	OS_MINIX    = "MINIX"
	OS_IRIX     = "IRIX"
	OS_FreeMiNT = "FreeMiNT"
	OS_Bitrig   = "Bitrig"
	OS_CYGWIN   = "CYGWIN"
	OS_MSYS     = "MSYS"
	OS_MINGW    = "MINGW"

	// Dead end
	OS_Uknown = "Uknown"
)

var supportedOses = []OSName{
	OS_AIX,
	OS_Android,
	OS_Darwin,
	OS_Dragonfly,
	OS_FreeBSD,
	OS_HURD,
	OS_Illumos,
	OS_JS,
	OS_Linux,
	OS_NACL,
	OS_NetBSD,
	OS_OpenBSD,
	OS_Plan9,
	OS_Solaris,
	OS_Windows,
	OS_ZOS,

	OS_MacOSX,
	OS_iOS,
	OS_OS2,
	OS_Haiku,
	OS_BeOS,
	OS_MINIX,
	OS_IRIX,
	OS_FreeMiNT,
	OS_Bitrig,
	OS_CYGWIN,
	OS_MSYS,
	OS_MINGW,
}
