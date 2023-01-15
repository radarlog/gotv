package main

import (
	"flag"
	"fmt"
	"os"
)

type GoTv []Channel

type Channel struct {
	Name      string
	Title     string
	StreamUrl string
	LogoUrl   string
}

func main() {
	configFile := flag.String("config", relativePath("config.yml"), "config file to read configuration from")
	m3uFile := flag.String("m3u", relativePath("gotv.m3u"), "m3u file to save a new playlist into")
	flag.Parse()

	config := load(*configFile)

	gotv := config.parse()

	count := gotv.save(*m3uFile)
	fmt.Printf("%d channels were successfully saved to %s \n", count, *m3uFile)

	os.Exit(0)
}
