package main

import (
	"log"
	"os"
	"strconv"
	"testing"
)

var (
	// Hardcoded for testing
	// values taken from result of wc filename
	filename string = "text.txt"
	dataByteSize int = 342187
	dataWords int = 58164
	dataLines int = 7145
)

func TestGetBytes(t *testing.T) {
	if GetBytes(filename) != int64(dataByteSize) {
		t.Error("Bytes does not match wc result of", dataByteSize)
	}
}

func TestGetWords(t *testing.T) {
	if GetWords(filename) != dataWords {
		t.Error("Words does not match wc result of", dataWords)
	}
}

func TestGetLines(t *testing.T) {
	if GetLines(filename) != dataLines {
		t.Error("Lines does not match wc result of", dataLines)
	}
}

func TestPipedInputGetDataOptimal(t *testing.T) {
	bs, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	
	result := strconv.Itoa(dataLines) + " " + strconv.Itoa(dataWords) + " " + strconv.Itoa(dataByteSize)
	if PipedInputGetDataOptimal(bs) != result {
		t.Error("Results does not match wc result of", result)
	}
}