# ![](https://fonts.gstatic.com/s/i/materialiconsoutlined/flare/v4/24px.svg) Detailed OS Detection at Runtime
[![](https://img.shields.io/github/v/release/codemodify/systemkit-platform-os?style=flat-square)](https://github.com/codemodify/systemkit-platform-os/releases/latest)
![](https://img.shields.io/github/languages/code-size/codemodify/systemkit-platform-os?style=flat-square)
![](https://img.shields.io/github/last-commit/codemodify/systemkit-platform-os?style=flat-square)
[![](https://img.shields.io/badge/license-0--license-brightgreen?style=flat-square)](https://github.com/codemodify/TheFreeLicense)

![](https://img.shields.io/github/workflow/status/codemodify/systemkit-platform-os/qa?style=flat-square)
![](https://img.shields.io/github/issues/codemodify/systemkit-platform-os?style=flat-square)
[![](https://goreportcard.com/badge/github.com/codemodify/systemkit-platform-os?style=flat-square)](https://goreportcard.com/report/github.com/codemodify/systemkit-platform-os)

[![](https://img.shields.io/badge/godoc-reference-brightgreen?style=flat-square)](https://godoc.org/github.com/codemodify/systemkit-platform-os)
![](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)
![](https://img.shields.io/gitter/room/codemodify/systemkit-platform-os?style=flat-square)

![](https://img.shields.io/github/contributors/codemodify/systemkit-platform-os?style=flat-square)
![](https://img.shields.io/github/stars/codemodify/systemkit-platform-os?style=flat-square)
![](https://img.shields.io/github/watchers/codemodify/systemkit-platform-os?style=flat-square)
![](https://img.shields.io/github/forks/codemodify/systemkit-platform-os?style=flat-square)

# ![](https://fonts.gstatic.com/s/i/materialicons/bookmarks/v4/24px.svg) Usage
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
	![](https://raw.githubusercontent.com/codemodify/systemkit-platform-os/master/.helper-files/dox/samples.png)
