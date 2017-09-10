package main

import (
	"fmt"

	"github.com/radarlog/gotv/onelike"
)

func main() {
	tv := onelike.Meta("src/github.com/radarlog/gotv/onelike/meta.yml")
	channels := tv.Channels()

	fmt.Println(channels)
}
