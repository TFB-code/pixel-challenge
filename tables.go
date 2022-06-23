package main

import (
	"fmt"
	"io/fs"
)

func populateTable(list []fs.FileInfo) []tableEntry {

	tableLength := len(list)
	table := make([]tableEntry, tableLength)

	for i, v := range list {
		table[i] = tableEntry{v.Name(), nil, 0, 0}
	}

	return table
}

func findMatches(table []tableEntry, pixel1 pixel) {
	for entry := range table {
		pixel2 := readPixel(table[entry].filePointer)
		if pixel1 == pixel2 {
			table[entry].matches = table[entry].matches + 1
		}
	}
}

func calculatePercentages(table []tableEntry, numberOfPixels float64) {
	for entry := range table {
		table[entry].matchPercentage = (table[entry].matches / numberOfPixels) * 100
	}
}

func outputResults(table []tableEntry) {
	fmt.Println()
	for entry := range table {
		filename := table[entry].filename
		fmt.Printf("  %s     %v matches       %06.3f%%\n", filename[len(filename)-16:],
			table[entry].matches,
			table[entry].matchPercentage)
	}
}
