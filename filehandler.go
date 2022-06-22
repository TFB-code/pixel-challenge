package main

import "os"
import "log"

func parseCommandLineArgs() (string, string) {
	argumentCount := len(os.Args) - 1
	if argumentCount < 1 {
		log.Fatalf("error - missing command line arguments\ncommand format is :\n   pixchallenge <directory> <filename>  or  pixchallenge <filename>")
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
