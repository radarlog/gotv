package main

import (
	"log"

	onlytv "github.com/radarlog/gotv/plugins"
)

func (c *config) parse() (channels GoTv) {
	if len(c.Channels) == 0 {
		log.Fatal("config: No `channels` have been loaded")
	}

	for _, item := range c.Channels {
		channel := item.Channel

		channels = append(channels, Channel{
			Name:      item.Name,
			Title:     channel.Title,
			StreamUrl: channel.findStream(),
			LogoUrl:   channel.LogoUrl,
		})
	}

	return channels
}

func (c *channel) findStream() (streamUrl string) {
	switch c.Plugin {
	case "onlytv":
		streamUrl = onlytv.FindStream(c.PageUrl)
	default:
		log.Fatalf("config: Channel %s has invalid `source`", c.Title)
	}

	return
}
