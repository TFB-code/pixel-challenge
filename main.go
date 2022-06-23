package main

import (
	"fmt"
)

func main() {

	searchdir, filename := parseCommandLineArgs()

	referenceFile := open(searchdir + filename)
	defer referenceFile.Close()

	directory := open(searchdir)
	defer directory.Close()

	filelist := makeFilelist(directory)
	filteredFilelist := filterFilelist(filelist, referenceFile)

	pix := readPixel(referenceFile)
	for _, v := range filteredFilelist {
		fmt.Println(v.Name(), " pixels ", v.Size()/3, " filesize ", v.Size())
	}
	fmt.Println()
	fmt.Println(pix)
}
