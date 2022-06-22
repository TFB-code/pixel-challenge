package main

import (
	"fmt"
	"log"
)

func main() {
	var referenceImage image
	searchdir, filename := parseCommandLineArgs()
	referenceFile := openReferenceFile(searchdir + filename)
	defer referenceFile.Close()

	pix, err := readPixel(referenceFile)
	if err != nil {
		log.Fatal(err)
	}
	referenceImage.pixel = append(referenceImage.pixel, pix)
	fmt.Println(referenceImage)

	// TODO read a byte.

}
