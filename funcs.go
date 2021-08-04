package main

import (
	"encoding/csv"
	"flag"
	"io"
	"log"
	"net/url"
	"strings"
)

// getInOutFileNames gets input and output file names via arguments -i and -o.
func getInOutFileNames(inFile, outFile *string) {
	flag.StringVar(
		inFile,
		"i",
		"in.csv",
		"path to csv file with logins exported from Bitwarden",
	)

	flag.StringVar(
		outFile,
		"o",
		"out.csv",
		"csv file name which will has logins in format accepted by 1Password",
	)
	flag.Parse()
}

// appendCsvToOutFileName checks if output filename has ".csv" at the end and if not adds it.
func appendCsvToOutFileName(outFile *string) {
	if !strings.HasSuffix(*outFile, ".csv") {
		*outFile += ".csv"
	}
}

// convert converts data between Bitwarden and 1Password csv format.
func convert(in io.Reader, out io.Writer) {
	r := csv.NewReader(in)
	w := csv.NewWriter(out)
	w.Write([]string{"title", "website", "username", "password", "notes"})
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		handleFatalErr(err)
		for _, val := range record {
			if val == "login" {
				seerviceURL, err := url.Parse(record[7])
				handleErr(err)
				w.Write([]string{record[3], seerviceURL.Host, record[8], record[9], record[4]})
				handleFatalErr(w.Error())
			}
		}
	}
	w.Flush()
	handleFatalErr(w.Error())
}

func handleErr(err error) {
	if err != nil {
		log.Println(err)
	}
}

func handleFatalErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
