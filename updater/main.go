package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	metaFile := flag.String("meta", "meta.yml", "meta file to read configuration from")
	dumpFile := flag.String("dump", "gotv.m3u", "m3u file to dump a new playlist into")
	flag.Parse()

	config := parse(*metaFile)

	count := config.dump(*dumpFile)
	fmt.Printf("%d channels were successfully dumped to %s \n", count, *dumpFile)

	os.Exit(0)
}
