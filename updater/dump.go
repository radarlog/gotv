package main

import (
	"fmt"
	"os"
	"text/template"
)

const tmpl = `#EXTM3U
{{range .}}
#EXTINF:0 tvg-logo="{{ .LogoUrl }}",{{ .Name }}
{{ .PageUrl }}
{{end}}`

func (config *meta) dump(file string) error {
	t, err := template.New("playlist").Parse(tmpl)
	if err != nil {
		return err
	}

	// Open a new file for writing only
	f, err := os.OpenFile(file, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	fmt.Println(file)

	return t.Execute(f, config.Channels)
}
