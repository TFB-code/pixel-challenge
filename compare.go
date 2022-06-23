package main

import "io/fs"

func populateTable(list []fs.FileInfo) []tableEntry {

	tableLength := len(list)
	table := make([]tableEntry, tableLength)

	for i, v := range list {
		table[i] = tableEntry{v.Name(), 0, 0}
	}

	return table
}

func compare() {}
