package main

import (
	"fmt"
	"os"
	cert "training/certificat_generator/certificate"
	"training/certificat_generator/html"
)

func main() {
	c, err := cert.New("Golang programming", "Bob Dylan", "2018-06-21")
	if err != nil {
		fmt.Printf("Error during the creation of the certificate : '%v'", err)
		os.Exit(1)
	}

	var saver cert.Saver
	//saver, err = pdf.New("output")
	saver, err = html.New("output")
	if err != nil {
		fmt.Println("Impossible to load the pdf generator")
		os.Exit(1)
	}
	saver.Save(*c)
}
