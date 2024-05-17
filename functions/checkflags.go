package ascciartfs

import (
	"fmt"
	"os"
	"strings"
)

func CheckFlag(arg string) string {
	output := "--output="
	if !strings.HasPrefix(arg, output) {
		return ""
	}
	file_name := arg[len(output):]
	if !strings.HasSuffix(file_name, ".txt") {
		fmt.Println("Error message")
		os.Exit(0)
	}
	return file_name
}
