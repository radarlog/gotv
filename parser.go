package main

import (
	"log"

	teliklive "github.com/radarlog/gotv/plugins/teliklive"
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
	case "teliklive":
		streamUrl = teliklive.FindStream(c.PageUrl)
	default:
		log.Fatalf("config: Channel %s has unknown `source`", c.Title)
	}

	return
}
