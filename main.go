package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	arglen := len(os.Args)
	if arglen < 2 {
		log.Fatal("No arguments specified")
	}

	file := os.Args[1]

	f, err := os.Stat(file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(f.Size(), file)
}
