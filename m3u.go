package main

import (
	"log"
	"os"
	"text/template"
)

// m3u file template
const tmpl = `#EXTM3U
{{range .}}
#EXTINF:0 tvg-logo="{{ .LogoUrl }}",{{ .Title }}
{{ .StreamUrl }}
{{end}}`

// m3u the config as a m3u file and return count of successfully processed channels
func (g GoTv) save(file string) int {
	t, err := template.New("playlist").Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}

	// Open a new file for writing only
	f, err := os.OpenFile(file, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}

	if err = t.Execute(f, g); err != nil {
		log.Fatal(err)
	}

	if err = f.Close(); err != nil {
		log.Fatal(err)
	}

	return len(g)
}
