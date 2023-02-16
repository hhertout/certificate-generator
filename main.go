package main

import (
	"flag"
	"fmt"
	"os"
	cert "training/certificat_generator/certificate"
	"training/certificat_generator/html"
	"training/certificat_generator/pdf"
)

func main() {
	outputType := flag.String("type", "pdf", "Output type of the certificate.")
	file := flag.String("file", "", "CSV file input")

	flag.Parse()

	if len(*file) <= 0 {
		fmt.Printf("Invalid file. Got=%v", *file)
	}

	var saver cert.Saver
	var err error

	switch *outputType {
	case "html":
		saver, err = html.New("output")
	case "pdf":
		saver, err = pdf.New("output")
	default:
		fmt.Println("Unknow output type.")
	}
	if err != nil {
		fmt.Println("Error during the generation of your certificate")
		os.Exit(1)
	}

	certs, err := cert.ParseCSV(*file)
	if err != nil {
		fmt.Printf("Could not parse the CSV file: %v", err)
		os.Exit(1)
	}

	for _, c := range certs {
		err = saver.Save(*c)
		if err != nil {
			fmt.Printf("Could not save Cert. (got=%v)", err)
		}
	}
}
