package main

func main() {
	searchdir, filename := parseCommandLineArgs()
	referenceFile := openReferenceFile(searchdir + filename)
	defer referenceFile.Close()

	// TODO read a byte.

}
