package main

import (
	"flag"
	"os"
)

func main() {
	metaFile := flag.String("meta", "meta.yml", "meta file to read configuration from")
	dumpFile := flag.String("dump", "", "m3u file to dump a new playlist into")
	flag.Parse()

	config := parse(*metaFile)

	if *dumpFile != "" {
		config.dump(*dumpFile)
		os.Exit(0)
	}

	config.build()
}
