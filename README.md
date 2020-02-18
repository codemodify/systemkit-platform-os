# Detailed OS Detection at Runtime

# Usage
- as a library in Go code
	```go
	package main

	import (
		"encoding/json"
		"fmt"
		"os"

		platformOS "github.com/codemodify/systemkit-platform-os"
	)

	func main() {
		data, err := json.Marshal(platformOS.Info())
		if err != nil {
			fmt.Println(err.Error())
		}

		if len(os.Args) > 1 && os.Args[1] == "-p" {
			data, err = json.MarshalIndent(platformOS.Info(), "", "    ")
			if err != nil {
				fmt.Println(err.Error())
			}
		}

		if

		fmt.Println(string(data))

		if osInfo.Name == platformOS.OS_Linux {

			fmt.Println("WE ARE ON Linux")
			fmt.Println("Distro: ", osInfo.Distribution)

		} else if platformOS.IsBSD(osInfo.Name) {

			fmt.Println("WE ARE ON BSD family")

		} else if platformOS.IsDarwin(osInfo.Name) {

			fmt.Println("WE ARE ON WINDOWS family")

		} else if platformOS.IsWindows(osInfo.Name) {

			fmt.Println("WE ARE ON WINDOWS family")
		}
	}
	```

- as a binary on a a bunch of platforms
	- `https://github.com/codemodify/systemkit-platform-os/releases/latest`
	![](https://raw.githubusercontent.com/codemodify/systemkit-platform-os/master/.helper-files/dox/sample.png)
