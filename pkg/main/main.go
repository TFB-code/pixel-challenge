package main

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

	openTabulatedFiles(table, searchdir)
	defer closeTabulatedFiles(table)

	for count := float64(0); count < numberOfPixels; count = count + 1 {
		pixel1 := readPixel(referenceFile)
		findMatches(table, pixel1)
	}

	calculatePercentages(table, numberOfPixels)

	outputResults(table)

}
