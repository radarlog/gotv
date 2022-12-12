package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	configFile := flag.String("config", relativePath("config.yml"), "config file to read configuration from")
	m3uFile := flag.String("m3u", relativePath("gotv.m3u"), "m3u file to save a new playlist into")
	flag.Parse()

	config := load(*configFile)

	count := config.save(*m3uFile)
	fmt.Printf("%d channels were successfully saved to %s \n", count, *m3uFile)

	os.Exit(0)
}
