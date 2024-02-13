package main

import (
	"fmt"
	"io"
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

		fileInfo, err := os.Stdin.Stat()
		if err != nil {
			log.Fatal(err)
		}

		if fileInfo.Mode()&os.ModeCharDevice != 0 {
			fmt.Println("No data piped to stdin")
			fmt.Println("No arguments specified")
		} else {
			bytes, err := io.ReadAll(os.Stdin)
			if err != nil {
				log.Fatal(err)
			}
			pipedInputGetData(bytes)
		}

	} else if arglen == 2 {
		filename = args[1]
		data := getAllData(filename)
		lines, words, ch := data[0], data[1], data[2]
		fmt.Println(lines, words, ch, filename)
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

// This funtion will return the number of lines, words, characters and filename
func getAllData(filename string) []int {
	ch := getBytes(filename)
	words := getWords(filename)
	lines := getLines(filename)

	return []int{lines, words, int(ch)}
}

// This function will take the stream of text data which is piped to our program to stdin
// and prints the number of lines, words and characters in the stream of text data

// We could have overloaded the function getAllData such that it could take a bytestream,
// But Go does not support function overloading (atleas as of now) and
// the logic needs to be changed since the getAllData works on file and we want to work on
// bytestream
func pipedInputGetData(bytestream []byte) {
	// we are writing the data to a temporary file, then calling the getAllData function to get
	// the results, and finally deleting the temporary file.
	// this is not optimal, and another way of doing it is to convert the bytestream to string
	// and perform the strings.split("\n") and strings.Fields(string) to get lines and words
	// but still needs to come up with logic for chars and bytes since we currently are using f.stat()
	// which uses a file pointer
	err := os.WriteFile("temp.txt", bytestream, 0666)
	if err != nil {
		log.Fatal(err)
	}
	data := getAllData("temp.txt")
	lines, words, ch := data[0], data[1], data[2]
	fmt.Println(lines, words, ch)


	err = os.Remove("temp.txt")
	if err != nil {
		log.Fatal(err)
	}
}