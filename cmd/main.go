package main

import (
	"encoding/json"
	"fmt"
	"os"

	platformOS "github.com/codemodify/systemkit-platform-os"
)

func main() {
	osInfo := platformOS.Info()

	data, err := json.Marshal(osInfo)
	if err != nil {
		fmt.Println(err.Error())
	}

	if len(os.Args) > 1 && os.Args[1] == "-p" {
		data, err = json.MarshalIndent(osInfo, "", "    ")
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	fmt.Println(string(data))
}
