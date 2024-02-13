package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	arglen := len(os.Args)
	args := os.Args
	var flag string
	var filename string

	if arglen == 1 {
		log.Fatal("No arguments specified")
	} else if arglen == 2 {
		filename = args[1]
		// Todo: needs to implement wc filename
		// output would be lines, words, chars, filename
	} else if arglen == 3 {
		flag = args[1]
		filename = args[2]
	
		switch flag {
		case "-c":
			getBytes(filename)
		case "-w":
			// getWords(filename)
		case "-l":
			getLines(filename)
		case "-m":
			getBytes(filename)
		default:
			log.Fatal("Flag provided but not defined")
		}		
	}	
}

func getBytes(filename string) {
	f, err := os.Stat(filename)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(f.Size(), filename)
}

func getLines(filename string) {
	f, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	
	stringf := string(f)
	lines := strings.Split(stringf, "\n")

	fmt.Println(len(lines), filename)
}