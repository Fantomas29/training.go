package main

import (
	"flag"
	"fmt"
	"os"

	"training.go/gencert/cert"
	"training.go/gencert/cert/html"
	"training.go/gencert/pdf"
)

func main() {
	file := flag.String("file", "", "CSV file input")
	outputType := flag.String("type", "pdf", "Output type of the certificate")
	flag.Parse()

	if len(*file) <= 0 {
		fmt.Printf("Invalid file. got=%v\n", *file)
		os.Exit(1)
	}

	var saver cert.Saver
	var err error
	switch *outputType {
	case "html":
		saver, err = html.New("output")
	case "pdf":
		saver, err = pdf.New("output")
	default:
		fmt.Printf("Unknown output type. got=%v", err)
	}

	if err != nil {
		fmt.Printf("Could not create generator: %v", err)
		os.Exit(1)
	}
	certs, err := cert.ParseCSV(*file)
	if err != nil {
		fmt.Printf("Could not parse CSV file: %v", err)
		os.Exit(1)
	}

	for _, c := range certs {
		err = saver.Save(*c)
		if err != nil {
			fmt.Printf("Could not save Cert. got=%v\n", err)
		}
	}
}
