package onelike

import (
	"errors"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Channel struct {
	Name      string `yaml:"name"`
	PageUrl   string `yaml:"page_url"`
	LogoUrl   string `yaml:"logo_url"`
	StreamUrl string
}

type Tv struct {
	HostUrl string `yaml:"host_url"`
	Channel map[string]*Channel `yaml:"channels"`
}

func Meta(metaFile string) (tv Tv) {
	data, err := ioutil.ReadFile(metaFile)
	if err != nil {
		log.Fatal(err)
	}

	if err := tv.parse(data); err != nil {
		log.Fatal(err)
	}

	return tv
}

func (tv *Tv) parse(data []byte) error {
	if err := yaml.Unmarshal(data, tv); err != nil {
		return err
	}

	if tv.HostUrl == "" {
		return errors.New("TV meta: `HostUrl` cannot be empty")
	}

	if len(tv.Channel) == 0 {
		return errors.New("TV meta: No channels found")
	}

	for name, channel := range tv.Channel {
		if channel.Name == "" {
			return errors.New("Channel " + name + " meta: `Name` cannot be empty")
		}

		if channel.PageUrl != "" {
			channel.PageUrl = tv.HostUrl + channel.PageUrl
		} else {
			return errors.New("Channel " + name + " meta: `PageUrl` cannot be empty")
		}

		if channel.LogoUrl != "" {
			channel.LogoUrl = tv.HostUrl + channel.LogoUrl
		}
	}

	return nil
}
