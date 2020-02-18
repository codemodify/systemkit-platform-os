# Detailed OS Detection at Runtime

# Usage in code
- Intented use is as a library in Go code
	```go
	package main

	import platform "github.com/codemodify/systemkit-platform"

	func main() {

		// Example 1
		fmt.Println(platform.GetInfo().CPU.Architecture)     // => ex: arm OR amd64
		fmt.Println(platform.GetInfo().CPU.Variant.Name)     // => ex: armv5 OR armv6 OR armv8, etc
		fmt.Println(platform.GetInfo().CPU.Variant.Detailed) // => ex: armv5te

		// Example 2
		if platform.IsArm(platform.GetInfo().CPU.Architecture) {
			if if platform.GetInfo().CPU.Variant.Name == platform.CPUV_ARMv8 {
				// we got AMRv8
			} else if platform.GetInfo().CPU.Variant.Name == platform.CPUV_ARMv5 {
				if platform.GetInfo().CPU.Variant.Detailed == platform.CPUVD_ARMv5T {
					// we got ARMv5t
				} else if platform.GetInfo().CPU.Variant.Detailed == platform.CPUVD_ARMv5TE {
					// we got ARMv5te
				} else if platform.GetInfo().CPU.Variant.Detailed == platform.CPUVD_ARMv5TEJ {
					// we got ARMv5tej
				}
			}
		}
	}
	```
