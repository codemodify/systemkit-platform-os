package os

const (
	OS_FreeBSD  Name = "FreeBSD"  // by GOOS - https://github.com/golang/go/blob/master/src/go/build/syslist.go
	OS_OrbisOS  Name = "Orbis OS" // https://en.wikipedia.org/wiki/PlayStation_4 (freebsd)
	OS_GhostBSD Name = "GhostBSD" //
	OS_NomadBSD Name = "NomadBSD" //
	OS_FuryBSD  Name = "FuryBSD"  //
)

var osNameFamilyFreeBSD = []Name{
	OS_FreeBSD,
	OS_OrbisOS,
	OS_GhostBSD,
	OS_NomadBSD,
	OS_FuryBSD,
}
