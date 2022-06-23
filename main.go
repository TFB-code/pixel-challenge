package main

import (
	"fmt"
	"log"
)

func main() {

	var referenceImage image
	searchdir, filename := parseCommandLineArgs()

	referenceFile := open(searchdir + filename)
	defer referenceFile.Close()

	directory := open(searchdir)
	defer directory.Close()

	filelist := filterFilelist(directory, referenceFile)

	pix, err := readPixel(referenceFile)
	if err != nil {
		log.Fatal(err)
	}
	referenceImage.pixel = append(referenceImage.pixel, pix)
	fmt.Println(referenceImage)

	for _, v := range filelist {
		fmt.Println(v.Name())
	}

}
