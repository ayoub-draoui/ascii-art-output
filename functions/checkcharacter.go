package ascciartfs

import (
	"fmt"
	"os"
	"strings"
)

func CheckCharacter(text string) []string {
	for _, char := range text {
		if rune(char) < 32 || rune(char) > 126 {
			fmt.Println("err:please enter a valid character")
			os.Exit(1)
		}
	}
	line := strings.Split(text, "\\n")
	return line
}
