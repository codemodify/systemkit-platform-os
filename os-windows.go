// +build windows

package os

import (
	"golang.org/x/sys/windows/registry"
)

func Info() OS {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows NT\CurrentVersion`, registry.QUERY_VALUE)
	if err != nil {
		return OS{
			Name: OS_Uknown,
		}
	}
	defer k.Close()

	version, _, err := k.GetStringValue("CurrentVersion")
	if err != nil {
		return OS{
			Name: OS_Uknown,
		}
	}

	productName, _, err := k.GetStringValue("ProductName")
	if err != nil {
		return OS{
			Name: OS_Uknown,
		}
	}

	return OS{
		Name:         OS_Windows,
		Version:      version,
		Distribution: productName,
		Features:     []string{},
	}
}
