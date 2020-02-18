package os

import (
	"bytes"
	"os/exec"
	"strings"
)

func runBinaryFetchOutput(binary string, args []string) ([]string, error) {
	var out bytes.Buffer
	var stderr bytes.Buffer

	cmd := exec.Command(binary, args...)
	cmd.Stdin = strings.NewReader("")
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return nil, err
	}

	return strings.Split(strings.TrimSpace(out.String()), "\n"), nil
}

func fetchFromLineToRight(lines []string, pattern string, split string) string {
	for _, line := range lines {
		list := strings.Split(line, split)
		if len(list) > 1 && strings.Contains(strings.TrimSpace(list[0]), pattern) {
			return strings.TrimSpace(list[1])
		}
	}

	return ""
}

// func fetchFromLineToLeft(lines []string, pattern string, split string) string {
// 	for _, line := range lines {
// 		list := strings.Split(line, split)
// 		if len(list) > 1 && strings.Contains(strings.TrimSpace(list[1]), pattern) {
// 			return strings.TrimSpace(list[0])
// 		}
// 	}
//
// 	return ""
// }

func fetchFirstInList(line string, split string) string {
	list := strings.Split(line, split)
	if len(list) > 0 {
		return strings.TrimSpace(list[0])
	}

	return ""
}

func fetchSecondInList(line string, split string) string {
	list := strings.Split(line, split)
	if len(list) > 1 {
		return strings.TrimSpace(list[1])
	}

	return ""
}

func cleanUp(value string, cleanItems []string) string {
	for _, item := range cleanItems {
		value = strings.ReplaceAll(value, item, "")
	}

	return value
}
