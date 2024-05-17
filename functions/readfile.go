package ascciartfs

import (
	"bufio"
	"log"
	"os"
)

func ReadFile(fileName string) map[rune][]string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalln("someting went wrong!")
	}
	defer file.Close()
	data := make(map[rune][]string)
	scaner := bufio.NewScanner(file)
	i := 32
	line := 0
	for scaner.Scan() {
		if scaner.Text() == "" {
			continue
		}
		if line == 8 {
			line = 0
			i++
		}
		data[rune(i)] = append(data[rune(i)], scaner.Text())
		line++
		if i == 127 {
			break
		}
	}
	return data
}
