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
	flag.Parse()

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

	c, err := cert.New("Golang programming", "Bob Dylan", "2018-06-21")
	if err != nil {
		fmt.Printf("Error during the creation of the certificate : '%v'", err)
		os.Exit(1)
	}

	saver.Save(*c)
}
