package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	ascciartfs "asciiartfs/functions"
)

func main() {
	readFile := map[rune][]string{}
	if len(os.Args) != 2 && len(os.Args) != 3 && len(os.Args) != 4 {
		log.Fatalln("err: you shoud enter two or three or foor arguments")
	}
	output := ascciartfs.CheckFlag(os.Args[1])
	args := os.Args[1:]
	banner := "standard"
	text := ""
	if len(args) == 3 {
		text = args[1]
		banner = args[2]
	} else if len(args) == 2 {
		text = args[0]
		banner = args[1]
	} else if len(args) == 1 {
		text = args[0]
	}
	if output != "" {
		if !strings.HasSuffix(banner, ".txt") {
			readFile = ascciartfs.ReadFile("./sources/" + banner + ".txt")
		} else {
			log.Fatalln("err: the third argument should be one of these file names (standerd),(shadow),(thinkertoy)")
		}
		// check if the input is among ascii manual
		checkcharacter := ascciartfs.CheckCharacter(text)
	
		// go print my argument
		result := ascciartfs.FindAndPrint(checkcharacter, readFile)
		file, err := os.Create(output)
		if err != nil {
			fmt.Println("err")
			os.Exit(0)
		}
		defer file.Close()
		file.WriteString(result)
	} else {
		if len(args) == 2 {
			readFile = ascciartfs.ReadFile("./sources/"+ banner +".txt")
			// if it's three arguments
		} else if !strings.HasSuffix(banner, ".txt") {
			readFile = ascciartfs.ReadFile("./sources/" + banner + ".txt")
		} else {
			log.Fatalln("err: the third argument should be one of these file names (standerd),(shadow),(thinkertoy)")
		}
		// check if the input is among ascii manual
	
		checkcharacter := ascciartfs.CheckCharacter(text)
	
		// go print my argument
		result := ascciartfs.FindAndPrint(checkcharacter, readFile)
		fmt.Print(result)
	}
}
