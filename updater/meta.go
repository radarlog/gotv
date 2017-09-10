package main

import (
	"errors"
	"io/ioutil"
	"log"

	"github.com/radarlog/gotv/onelike"
	"gopkg.in/yaml.v2"
)

type Tv struct {
	Onelike onelike.Tv `yaml:"onelike"`
}

type Yaml struct {
	HostUrl string `yaml:"host_url"`
	LogoDir string `yaml:"logo_dir"`
	Tv      Tv `yaml:"tv_list"`
}

func Meta(metaFile string) (config Yaml) {
	data, err := ioutil.ReadFile(metaFile)
	if err != nil {
		log.Fatal(err)
	}

	if err := config.parse(data); err != nil {
		log.Fatal(err)
	}

	return config
}

func (config *Yaml) parse(data []byte) error {
	if err := yaml.UnmarshalStrict(data, config); err != nil {
		return err
	}

	if config.HostUrl == "" {
		return errors.New("meta: `HostUrl` cannot be empty")
	}

	if config.LogoDir == "" {
		return errors.New("meta: `LogoDir` cannot be empty")
	}

	//if len(config.Tv) == 0 {
	//	return errors.New("TV meta: No channels found")
	//}

	return nil
}
