package main

import "os"

type pixel struct {
	red, green, blue int
}

type tableEntry struct {
	filename        string
	filePointer     *os.File
	matches         float64
	matchPercentage float64
}
