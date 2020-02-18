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

	fmt.Println(string(data))
}
