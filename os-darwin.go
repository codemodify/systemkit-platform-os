// +build darwin

package os

import "golang.org/x/sys/unix"

func info() OS {
	result := newOS()
	result.Family = Family_Darwin
	result.Name = OS_macOS

	output := runBinaryFetchOutput("sw_vers", []string{"-productVersion"})
	if len(output) > 0 {
		result.Version = output[0]
	}

	result.Kernel.Name = Kernel_XNU
	result.Kernel.Version, _ = unix.Sysctl("kern.osrelease")
	result.Kernel.BuildTime, _ = unix.Sysctl("kern.version")
	result.Kernel.ExecFormat = ExecFormat_MachO

	return result
}
