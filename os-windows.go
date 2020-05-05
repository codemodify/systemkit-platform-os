// +build windows

package os

import (
	"golang.org/x/sys/windows/registry"
)

func info() OS {
	result := newOS()
	result.Family = Family_Windows
	result.Name = OS_Windows

	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows NT\CurrentVersion`, registry.QUERY_VALUE)
	if err != nil {
		return result
	}
	defer k.Close()

	productName, _, err := k.GetStringValue("ProductName")
	if err == nil {
		result.Version = productName
	}

	output := runBinaryFetchOutput("wmic", []string{"os", "get", "Version"}) //  Caption
	if len(output) > 1 {
		result.Version = result.Version + ", Build: " + output[1]
	}

	kernelVersion, _, err := k.GetStringValue("CurrentVersion")
	if err == nil {
		result.Kernel.Name = "Windows NT"
		result.Kernel.Version = kernelVersion
	}

	kernelBuild, _, err := k.GetStringValue("BuildLab")
	if err == nil {
		result.Kernel.BuildTime = kernelBuild
	}

	result.Kernel.ExecFormat = ExecFormat_PE

	return result
}
