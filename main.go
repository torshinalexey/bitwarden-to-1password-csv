package main

import (
	"os"
)

func main() {
	var (
		inFile  string
		outFile string
	)

	getInOutFileNames(&inFile, &outFile)

	in, err := os.Open(inFile)
	handleFatalErr(err)
	defer in.Close()

	appendCsvToOutFileName(&outFile)
	out, err := os.Create(outFile)
	handleFatalErr(err)

	convert(in, out)
}
