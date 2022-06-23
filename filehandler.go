package main

import (
	"io/fs"
	"log"
	"os"
	"strings"
)

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

func open(name string) *os.File {
	pointer, err := os.Open(name)
	if err != nil {
		log.Fatalf("error - unable to find file or directory [%s]", name)
	}
	return pointer
}

func readPixel(file *os.File) pixel {

	var pixel pixel
	pixelBytes := make([]byte, 3)
	readCounter, err := file.Read(pixelBytes)
	if err != nil {
		log.Fatal(err)
	}

	if readCounter < 3 {
		log.Fatalf("error - unexpected end of file %v", file.Name())
	}

	pixel.red = int(pixelBytes[0])
	pixel.green = int(pixelBytes[1])
	pixel.blue = int(pixelBytes[2])

	return pixel
}

func makeFilelist(directory *os.File) []fs.DirEntry {

	filelist, err := directory.ReadDir(0)
	if err != nil {
		log.Fatal(err)
	}
	return filelist
}

func filterFilelist(filelist []fs.DirEntry, exclude *os.File) []fs.FileInfo {

	exclusionFileInfo, _ := exclude.Stat()

	var filteredList []fs.FileInfo
	for _, v := range filelist {

		currentFileName := strings.ToLower(v.Name())
		if v.IsDir() || currentFileName[len(currentFileName)-3:] != "raw" {
			continue
		}
		currentFileInfo, _ := v.Info()
		if !os.SameFile(currentFileInfo, exclusionFileInfo) {
			filteredList = append(filteredList, currentFileInfo)
		}
	}
	return filteredList
}
