package os

const (
	OS_Darwin     Name = "Darwin"     // by GOOS - https://github.com/golang/go/blob/master/src/go/build/syslist.go
	OS_PureDarwin Name = "PureDarwin" //
	OS_MacOSX     Name = "Mac OS X"   // Mac OS X 10.0 released in March 2001
	OS_macOS      Name = "macOS"      // "macOS" since 2016
	OS_iOS        Name = "iOS"        //
	OS_watchOS    Name = "watchOS"    //
	OS_tvOS       Name = "tvOS"       //
	OS_iPadOS     Name = "iPadOS"     //
)

var osNameFamilyDarwin = []Name{
	OS_Darwin,
	OS_PureDarwin,
	OS_MacOSX,
	OS_macOS,
	OS_iOS,
	OS_watchOS,
	OS_tvOS,
	OS_iPadOS,
}
