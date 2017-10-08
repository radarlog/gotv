package main

import "fmt"

func (config *meta) build() error {
	fmt.Println(config.Channels["mezzo"].PageUrl)
	fmt.Println(config.Channels["nat-geo-wild"].PageUrl)

	fmt.Println("==================================")

	for name, channel := range config.Channels {
		fmt.Println(name)
		fmt.Println(channel)
	}

	return nil
}
