package os

const (
	OS_Windows Name = "Windows"  // by GOOS - https://github.com/golang/go/blob/master/src/go/build/syslist.go
	OS_CYGWIN  Name = "CYGWIN"   //
	OS_MSYS    Name = "MSYS"     //
	OS_MINGW   Name = "MINGW"    //
	OS_ReactOS Name = "React OS" //
)

var osNameFamilyWindows = []Name{
	OS_Windows,
	OS_CYGWIN,
	OS_MSYS,
	OS_MINGW,
	OS_ReactOS,
}
