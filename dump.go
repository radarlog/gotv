package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"text/template"
)

type dump struct {
	Logo   string
	Name   string
	Stream string
}

const tmpl = `#EXTM3U
{{range .}}
#EXTINF:0 tvg-logo="{{ .Logo }}",{{ .Name }}
{{ .Stream }}
{{end}}`

func (config *meta) dump(file string) int {
	t, err := template.New("playlist").Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}

	// Open a new file for writing only
	f, err := os.OpenFile(file, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	dumpList := make([]dump, 0)
	for name, channel := range config.Channels {
		if channel.PageUrl != "" {
			logo, err := channel.dumpLogo(name, config.LogoDir)
			if err != nil {
				log.Fatal(err)
			}

			dumpList = append(dumpList, dump{
				Logo:   logo,
				Name:   channel.Name,
				Stream: channel.PageUrl,
			})
		}
	}

	err = t.Execute(f, dumpList)
	if err != nil {
		log.Fatal(err)
	}

	return len(dumpList)
}

func (c *Channel) dumpLogo(name string, dir string) (path string, err error) {
	// create logo dir
	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return
	}

	// file name
	path = fmt.Sprintf("%s/%s%s", dir, name, filepath.Ext(c.LogoUrl))

	// create logo file
	file, err := os.Create(path)
	if err != nil {
		return
	}
	defer file.Close()

	// download logo
	response, err := http.Get(c.LogoUrl)
	defer response.Body.Close()

	if response.StatusCode >= 200 && response.StatusCode <= 299 {
		// write downloaded logo to created file
		_, err = io.Copy(file, response.Body)
	}

	return
}