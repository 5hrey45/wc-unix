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
		getAllData(filename)
	} else if arglen == 3 {
		flag = args[1]
		filename = args[2]

		switch flag {
		case "-c":
			fmt.Println(getBytes(filename), filename)
		case "-w":
			fmt.Println(getWords(filename), filename)
		case "-l":
			fmt.Println(getLines(filename), filename)
		case "-m":
			fmt.Println(getBytes(filename), filename)
		default:
			log.Fatal("Flag provided but not defined")
		}
	}
}

// This funcion will return the number of bytes/characters in the file
func getBytes(filename string) int64 {
	f, err := os.Stat(filename)
	if err != nil {
		log.Fatal(err)
	}

	return f.Size()
}

// This function will return the number of lines in the file
func getLines(filename string) int {
	f := getFileByteStream(filename)

	// converting the bytestream f to string f
	stringf := string(f)
	// lines slice will store the lines obtained by splitting the string f by newline delimiter
	lines := strings.Split(stringf, "\n")

	// the size of the lines slice - 1 is the number of lines present in the file
	// why size -1 ?
	// the last line in many text file does not end with a newline char
	// the wc tool counts the number of \n (newline char) so in reality it's usually one line
	// short of the actual lines in the file
	return len(lines) - 1
}

// This function will return the number of words in the file
func getWords(filename string) int {
	f := getFileByteStream(filename)

	stringf := string(f)
	words := strings.Fields(stringf)
	// we cannot use strings.split(" ") since the words might be seperated by tabs or newlines as
	// word seperators
	// strings.Fields(s string) splits the string into words by taking all kinds of whitespace
	// seperators into account

	// the size of the words slice is the number of words present in the file
	return len(words)
}

// This funcion will return a bytestream
func getFileByteStream(filename string) []byte {
	// the os.ReadFile reads the file and returns a stream of bytes.
	f, err := os.ReadFile(filename)

	if err != nil {
		log.Fatal(err)
	}

	return f
}

// This funtion will print the number of lines, words, characters and filename
func getAllData(filename string) {
	ch := getBytes(filename)
	words := getWords(filename)
	lines := getLines(filename)

	fmt.Println(lines, words, ch, filename)
}
