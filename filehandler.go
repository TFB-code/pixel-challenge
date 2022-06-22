package main

import "os"
import "log"

func parseCommandLineArgs() (string, string) {
	argumentCount := len(os.Args) - 1
	if argumentCount < 1 {
		log.Fatalf("error - missing command line arguments\ncommand format is :\npixchallenge <directory> <filename>  or  pixchallenge <filename>")
	}
	if argumentCount == 1 {
		return "./", os.Args[1]
	}
	return os.Args[1], os.Args[2]
}

func openReferenceFile(filename string) *os.File {
	filePointer, err := os.Open(filename)
	if err != nil {
		log.Fatalf("error - unable to find file [%s]", filename)
	}
	return filePointer
}

func readPixel(file *os.File) (pixel, error) {
	var pixel pixel
	pixelBytes := make([]byte, 3)
	readCounter, err := file.Read(pixelBytes)
	if readCounter < 3 {
		log.Fatalf("error - unexpected end of file %v", file.Name())
	}
	pixel.red = int(pixelBytes[0])
	pixel.green = int(pixelBytes[1])
	pixel.blue = int(pixelBytes[2])
	return pixel, err
}
