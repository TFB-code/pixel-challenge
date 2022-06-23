package main

import (
	"fmt"
)

func main() {

	searchdir, filename := parseCommandLineArgs()

	referenceFile := open(searchdir + filename)
	defer referenceFile.Close()

	numberOfPixels := pixelCount(referenceFile)

	directory := open(searchdir)
	defer directory.Close()

	filelist := makeFilelist(directory)
	filteredFilelist := filterFilelist(filelist, referenceFile)

	table := populateTable(filteredFilelist)

	for entry := range table {
		table[entry].filePointer = open(searchdir + table[entry].filename)
		defer table[entry].filePointer.Close()
	}

	for count := float64(0); count < numberOfPixels; count = count + 1 {
		pixel := readPixel(referenceFile)
		for entry := range table {
			if pixel == readPixel(table[entry].filePointer) {
				table[entry].matches = table[entry].matches + 1
			}
		}
	}

	for entry := range table {
		table[entry].matchPercentage = (table[entry].matches / numberOfPixels) * 100
	}

	for i, v := range filteredFilelist {
		fmt.Println(v.Name(), " pixels ", v.Size()/3, " % ", table[i].matchPercentage)
	}

}
