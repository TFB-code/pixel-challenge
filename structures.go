package main

type pixel struct {
	red, green, blue int
}

type tableEntry struct {
	filename        string
	filePointer     int
	matchPercentage float32
}
