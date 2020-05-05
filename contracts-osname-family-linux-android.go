package os

const (
	OS_Android     Name = "Android"                // by GOOS - https://github.com/golang/go/blob/master/src/go/build/syslist.go
	OS_AndroidX86  Name = "Android x86"            //
	OS_CyanogenMod Name = "CyanogenMod/CyanogenOS" //
	OS_FireOS      Name = "Fire OS"                //
	OS_LineageOS   Name = "LineageOS"              //
)

var osNameFamilyLinuxAndroid = []Name{
	OS_Android,
	OS_AndroidX86,
	OS_CyanogenMod,
	OS_FireOS,
	OS_LineageOS,
}
