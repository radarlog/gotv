package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	onlytv "github.com/radarlog/gotv/plugins"
	"gopkg.in/yaml.v3"
)

// The representation of a config file
type config struct {
	Channels Channels `yaml:"channels"`
}

type Channels []Item

type Item struct {
	Name    string
	Channel Channel
}

// The representation of a channel in the config file
type Channel struct {
	Title   string `yaml:"title"`
	Plugin  string `yaml:"plugin"`
	PageUrl string `yaml:"page_url"`
	LogoUrl string `yaml:"logo_url"`
}

// custom marshaling to an intermediate yaml.Node
// https://stackoverflow.com/a/63260431
func (channels *Channels) UnmarshalYAML(value *yaml.Node) error {
	if value.Kind != yaml.MappingNode {
		return fmt.Errorf("Channel must contain YAML mapping, has %v", value.Kind)
	}

	*channels = make([]Item, len(value.Content)/2)

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

	config.validate()

	return
}

// perform config's validation and populate channel's stream by the corresponding source handler
// TODO: split validation and population into two different functions
func (config *config) validate() {
	if len(config.Channels) == 0 {
		log.Fatal("config: No `channels` have been found")
	}

	for _, item := range config.Channels {
		channel := item.Channel
		name := item.Name

		switch channel.Plugin {
		case "onlytv":
			channel.PageUrl = onlytv.FindStream(channel.PageUrl)
		default:
			log.Fatalf("config: Channel %s has invalid `source`", name)
		}
	}
}

func relativePath(p string) string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	path := filepath.Dir(ex)

	return fmt.Sprintf("%s/%s", path, p)
}
