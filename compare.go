package main

import (
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
