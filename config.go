package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// The representation of a config file
type config struct {
	Channels channels `yaml:"channels"`
}

type channels []item

type item struct {
	Name    string
	Channel channel
}

// The representation of a channel in the config file
type channel struct {
	Title   string `yaml:"title"`
	Plugin  string `yaml:"plugin"`
	PageUrl string `yaml:"page_url"`
	LogoUrl string `yaml:"logo_url"`
}

// custom marshaling to an intermediate yaml.Node
// https://stackoverflow.com/a/63260431
func (channels *channels) UnmarshalYAML(value *yaml.Node) error {
	if value.Kind != yaml.MappingNode {
		return fmt.Errorf("Channel must contain YAML mapping, has %v", value.Kind)
	}

	*channels = make([]item, len(value.Content)/2)

	for i := 0; i < len(value.Content); i += 2 {
		var res = &(*channels)[i/2]

		if err := value.Content[i].Decode(&res.Name); err != nil {
			return err
		}

		if err := value.Content[i+1].Decode(&res.Channel); err != nil {
			return err
		}
	}

	return nil
}

// load and parse a config file
func load(file string) (config config) {
	data, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	// parse yaml
	if err := yaml.Unmarshal(data, &config); err != nil {
		log.Fatal(err)
	}

	return
}

func relativePath(p string) string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	path := filepath.Dir(ex)

	return fmt.Sprintf("%s/%s", path, p)
}
