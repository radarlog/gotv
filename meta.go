package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/radarlog/gotv/plugins"
	"gopkg.in/yaml.v2"
)

// The representation of a config file
type meta struct {
	LogoDir  string              `yaml:"logo_dir"`
	Channels map[string]*Channel `yaml:"channels"`
}

// The representation of a channel in the config file
type Channel struct {
	Name    string `yaml:"name"`
	Plugin  string `yaml:"plugin"`
	PageUrl string `yaml:"page_url"`
	LogoUrl string `yaml:"logo_url"`
}

// load and parse a config file
func load(file string) (config meta) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	// parse yaml
	if err := yaml.UnmarshalStrict(data, &config); err != nil {
		log.Fatal(err)
	}

	if err := config.validate(); err != nil {
		log.Fatal(err)
	}

	return
}

// perform config's validation and populate channel's stream by the corresponding source handler
// TODO: split validation and population into two different functions
func (config *meta) validate() error {
	if config.LogoDir == "" {
		return errors.New("meta: `logo_dir` cannot be empty")
	}

	if len(config.Channels) == 0 {
		return errors.New("meta: No `channels` have been found")
	}

	for name, channel := range config.Channels {
		switch channel.Plugin {
		case "onelike":
			channel.PageUrl = onelike.FindStream(channel.PageUrl)
		case "":
			return errors.New(fmt.Sprintf("meta: Channel %s has no `source`", name))
		default:
			return errors.New(fmt.Sprintf("meta: Channel %s has invalid `source`", name))
		}
	}

	return nil
}
