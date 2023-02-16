package html

import (
	"fmt"
	"os"
	"path"
	"text/template"
	cert "training/certificat_generator/certificate"
)

type HtmlSaver struct {
	OutputDir string
}

func New(outputdir string) (*HtmlSaver, error) {
	var h *HtmlSaver
	err := os.MkdirAll(outputdir, os.ModePerm)
	if err != nil {
		return nil, err
	}
	h = &HtmlSaver{
		OutputDir: outputdir,
	}
	return h, nil
}

func (p *HtmlSaver) Save(cert cert.Certificate) error {
	t, err := template.New("certificate").Parse(tpl)
	if err != nil {
		return err
	}

	filename := fmt.Sprintf("%v.html", cert.LabelTitle)
	path := path.Join(p.OutputDir, filename)
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	err = t.Execute(f, cert)
	if err != nil {
		return err
	}

	fmt.Printf("HTML version of the certificate is saved in '%v'", path)

	return nil
}

var tpl = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{.LabelTitle}}</title>
	<style>
	body {
		text-align: center;
		font-family: 'Segoe UI', Tahoma, Geneva
	}
	h1 {
		font-size: 3em
	}
	</style>
</head>
<body>
    <h1>{{.LabelCompletion}}</h1>
    <h2>{{.LabelPresented}}</h2>
    <h1>{{.Name}}</h1>
	<h2>{{.LabelParticipation}}</h2>
	<p><em>{{.LabelDate}}</em></p>
</body>
</html>
`
