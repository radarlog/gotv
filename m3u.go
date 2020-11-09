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

// m3u file template
const tmpl = `#EXTM3U
{{range .}}
#EXTINF:0 tvg-logo="{{ .Logo }}",{{ .Name }}
{{ .Stream }}
{{end}}`

// m3u file representation
type m3u struct {
	Logo   string
	Name   string
	Stream string
}

// m3u the config as a m3u file and return count of successfully processed channels
func (config *config) save(file string) int {
	t, err := template.New("playlist").Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}

	// Open a new file for writing only
	f, err := os.OpenFile(file, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}

	saveList := make([]m3u, 0)
	for name, channel := range config.Channels {
		if channel.PageUrl != "" {
			logo, err := channel.saveLogo(name, config.LogoDir)
			if err != nil {
				log.Fatal(err)
			}

			saveList = append(saveList, m3u{
				Logo:   logo,
				Name:   channel.Name,
				Stream: channel.PageUrl,
			})
		}
	}

	if err = t.Execute(f, saveList); err != nil {
		log.Fatal(err)
	}

	if err = f.Close(); err != nil {
		log.Fatal(err)
	}

	return len(saveList)
}

// fetch and m3u channel's logo
func (c *Channel) saveLogo(name string, dir string) (path string, err error) {
	// create logo dir
	if err = os.MkdirAll(dir, os.ModePerm); err != nil {
		return
	}

	// file name
	path = fmt.Sprintf("%s/%s%s", dir, name, filepath.Ext(c.LogoUrl))

	// create a logo file
	file, err := os.Create(path)
	if err != nil {
		return
	}

	// download a given logo
	response, err := http.Get(c.LogoUrl)
	if err != nil {
		return
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		// write downloaded logo into the created file
		_, err = io.Copy(file, response.Body)

		if err = file.Close(); err != nil {
			return
		}
	}

	return
}
