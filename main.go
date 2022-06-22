package main

import "os"
import "log"

func openReferenceFile(filename string) *os.File {
	filePointer, err := os.Open(filename)
	if err != nil {
		log.Fatalf("error, unable to find file [%s]", filename)
	}
	return filePointer
}

func main() {

	referenceFile := openReferenceFile("abc")
	defer referenceFile.Close()

	// TODO read a byte.

}
