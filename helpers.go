package os

import (
	"os/exec"
	"strings"
)

func newOS() OS {
	return OS{
		Family:  Family_Unknown,
		Name:    OS_Unknown,
		Version: "",
		Kernel: KernelInfo{
			Name:       Kernel_Unknown,
			Version:    "",
			BuildTime:  "",
			ExecFormat: ExecFormat_Unknown,
		},
		Features: []string{},
	}
}

func runBinaryFetchOutput(binary string, args []string) []string {
	cmd := exec.Command(binary, args...)
	output, _ := cmd.CombinedOutput()
	return strings.Split(strings.TrimSpace(string(output)), "\n")
}

func fetchFirstInList(list []string) string {
	if len(list) > 0 {
		return strings.TrimSpace(list[0])
	}

	return ""
}

func fetchWordAt(manyWords string, split string, position int) string {
	list := strings.Split(manyWords, split)
	if len(list) > position {
		return removeInternalDoubleOrMoreSpaces(strings.TrimSpace(list[position]))
	}

	return ""
}

func stringInList(val string, values *[]string) bool {
	for i := 0; i < len(*values); i++ {
		if (*values)[i] == val {
			return true
		}
	}

	return false
}

func removeWords(value string, wordsToRemove []string) string {
	for _, word := range wordsToRemove {
		value = strings.ReplaceAll(value, word, "")
	}

	return removeInternalDoubleOrMoreSpaces(strings.TrimSpace(value))
}

func removeInternalDoubleOrMoreSpaces(val string) string {
	val = strings.ReplaceAll(val, "    ", "")
	val = strings.ReplaceAll(val, "   ", "")
	val = strings.ReplaceAll(val, "  ", "")

	return val
}
