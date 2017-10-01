package main

import (
	"flag"
	"fmt"
)

func main() {
	metaFile := flag.String("m", "meta.yml", "meta file to read configuration from")
	flag.Parse()

	config := Meta(*metaFile)

	fmt.Println(config.Tv.Onelike["mezzo"].Stream())
	fmt.Println(config.Tv.Onelike["nat-geo-wild"].Stream())

	fmt.Println("==================================")

	for name, channel := range config.Tv.Onelike {
		fmt.Println(name)
		fmt.Println(channel)
	}

	fmt.Println(config)
}
