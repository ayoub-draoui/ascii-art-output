package ascciartfs

func FindAndPrint(checkcharacter []string, readfile map[rune][]string) string {
	result := ""
	is_Printed := false
	for idx, word := range checkcharacter {
		if word != "" {
			i := 0
			for i < 8 {
				line := ""
				for _, char := range word {
					line = readfile[char][i]
					// fmt.Print(line)
					result += line
				}
				result += "\n"
				i++
				is_Printed = true
			}
		} else {
			if idx == len(checkcharacter)-1 && !is_Printed {
				continue
			} else {
				// fmt.Println()
				result += "\n"
			}
		}
	}
	return result
}
