package pdf

import (
	"fmt"
	"os"
	"path"
	cert "training/certificat_generator/certificate"

	"github.com/jung-kurt/gofpdf"
)

type PdfSaver struct {
	OutputDir string
}

func New(outputdir string) (*PdfSaver, error) {
	var p *PdfSaver
	err := os.MkdirAll(outputdir, os.ModePerm)
	if err != nil {
		return nil, err
	}
	p = &PdfSaver{
		OutputDir: outputdir,
	}
	return p, nil
}

func (p *PdfSaver) Save(cert cert.Certificate) error {
	pdf := gofpdf.New(gofpdf.OrientationLandscape, "mm", "A4", "")

	pdf.SetTitle(cert.LabelTitle, false)
	pdf.AddPage()

	// Creating the different parts of the pdf
	// Background
	background(pdf)

	//Header
	header(pdf, &cert)

	//Body
	body(pdf, &cert)

	//footer
	footer(pdf)

	filename := fmt.Sprintf("%v.pdf", cert.LabelTitle)
	path := path.Join(p.OutputDir, filename)
	err := pdf.OutputFileAndClose(path)
	if err != nil {
		return err
	}
	fmt.Printf("Process complete, certificate is save at '%v'\n", path)
	return nil
}

func background(pdf *gofpdf.Fpdf) {
	opts := gofpdf.ImageOptions{
		ImageType: "png",
	}
	pageWidth, pageHeight := pdf.GetPageSize()
	pdf.ImageOptions("assets/images/background.png", 0, 0, pageWidth, pageHeight, false, opts, 0, "")
}

func header(pdf *gofpdf.Fpdf, c *cert.Certificate) {
	// Place the left image
	margin := 30
	x := 0
	imageWidth := 30
	filename := "assets/images/gopher.jpg"

	opts := gofpdf.ImageOptions{
		ImageType: "jpg",
	}

	pdf.ImageOptions(filename, float64(x+margin), 20, float64(imageWidth), 0, false, opts, 0, "")

	// Place the right image
	pageWidth, _ := pdf.GetPageSize()
	x = int(pageWidth) - imageWidth
	pdf.ImageOptions(filename, float64(x-margin), 20, float64(imageWidth), 0, false, opts, 0, "")

	//Write the text
	pdf.SetFont("Helvetica", "", 40)
	pdf.WriteAligned(0, 50, c.LabelCompletion, "C")
	pdf.Ln(30)
}

func body(pdf *gofpdf.Fpdf, c *cert.Certificate) {
	pdf.SetFont("Helvetica", "B", 20)
	pdf.WriteAligned(0, 50, c.LabelPresented, "C")
	pdf.Ln(30)

	// Student-name
	pdf.SetFont("Times", "I", 40)
	pdf.WriteAligned(0, 50, c.Name, "C")
	pdf.Ln(30)

	// Participation
	pdf.SetFont("Helvetica", "B", 20)
	pdf.WriteAligned(0, 50, c.LabelParticipation, "C")
	pdf.Ln(30)

	// Date
	pdf.SetFont("Helvetica", "B", 15)
	pdf.WriteAligned(0, 50, c.LabelDate, "C")
}

func footer(pdf *gofpdf.Fpdf) {
	margin := float64(30)
	imageWidth := 30

	pageWidth, pageHeight := pdf.GetPageSize()
	x := pageWidth - float64(imageWidth)
	y := pageHeight - float64(imageWidth)

	filename := "assets/images/stamp.jpg"

	opts := gofpdf.ImageOptions{
		ImageType: "jpg",
	}

	pdf.ImageOptions(filename, float64(x-margin), float64(y-margin), float64(imageWidth), 0, false, opts, 0, "")
}
