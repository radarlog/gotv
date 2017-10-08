package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/radarlog/gotv/onelike"
	"gopkg.in/yaml.v2"
)

type Channel struct {
	Name    string `yaml:"name"`
	Handler string `yaml:"handler"`
	PageUrl string `yaml:"page_url"`
	LogoUrl string `yaml:"logo_url"`
}

type meta struct {
	HostUrl  string              `yaml:"host_url"`
	LogoDir  string              `yaml:"logo_dir"`
	Channels map[string]*Channel `yaml:"channels"`
}

func parse(file string) (config meta) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	if err := config.parse(data); err != nil {
		log.Fatal(err)
	}

	return config
}

func (config *meta) parse(data []byte) error {
	if err := yaml.UnmarshalStrict(data, config); err != nil {
		return err
	}

	if config.HostUrl == "" {
		return errors.New("meta: `host_url` cannot be empty")
	}

	if config.LogoDir == "" {
		return errors.New("meta: `logo_dir` cannot be empty")
	}

	if len(config.Channels) == 0 {
		return errors.New("meta: No `channels` have been found")
	}

	for name, channel := range config.Channels {
		switch channel.Handler {
		case "":
			return errors.New(fmt.Sprintf("meta: Channel %s has no `source`", name))
		case "onelike":
			channel.PageUrl = onelike.FindStream(channel.PageUrl)
		default:
			return errors.New(fmt.Sprintf("meta: Channel %s has invalid `source`", name))
		}
	}

	return nil
}
