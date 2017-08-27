package main

import (
	"fmt"

	"github.com/radarlog/iptv-wrapper/onelike"
)

func main() {
	streamUrl := onelike.Channel()

	fmt.Println(streamUrl)
}
