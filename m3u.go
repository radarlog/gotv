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

// directory where to save channel logos
const LogoDir = "logos/"

// m3u file template
const tmpl = `#EXTM3U
{{range .}}
#EXTINF:0 tvg-logo="{{ .Logo }}",{{ .Title }}
{{ .Stream }}
{{end}}`

// m3u file representation
type m3u struct {
	Logo   string
	Title  string
	Stream string
}

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

	saveList := make([]m3u, 0)
	for _, channel := range g {

		if channel.StreamUrl != "" {
			filename, err := channel.saveLogo(channel.Name)
			if err != nil {
				log.Fatal(err)
			}

			saveList = append(saveList, m3u{
				Logo:   filepath.Join(LogoDir, filename),
				Title:  channel.Title,
				Stream: channel.StreamUrl,
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
func (c *Channel) saveLogo(name string) (filename string, err error) {
	dir := relativePath(LogoDir)

	// create logo dir
	if err = os.MkdirAll(dir, os.ModePerm); err != nil {
		return
	}

	// file name
	filename = fmt.Sprintf("%s%s", name, filepath.Ext(c.LogoUrl))

	// create a logo file
	file, err := os.Create(filepath.Join(dir, filename))
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
