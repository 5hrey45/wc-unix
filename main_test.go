package main

import (
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
