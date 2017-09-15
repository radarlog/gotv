package main

import (
	"fmt"
)

func main() {
	config := Meta("src/github.com/radarlog/gotv/meta.yml")

	fmt.Println(config.Tv.Onelike["mezzo"].Stream())
	fmt.Println(config.Tv.Onelike["nat-geo-wild"].Stream())

	fmt.Println("==================================")

	for name, channel := range config.Tv.Onelike {
		fmt.Println(name)
		fmt.Println(channel)
	}

	fmt.Println(config)
}
