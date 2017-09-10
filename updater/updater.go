package main

import (
	"fmt"
)

func main() {
	config := Meta("src/github.com/radarlog/gotv/meta.yml")

	fmt.Println(config.Tv.Onelike.Mezzo.Stream())
	fmt.Println(config.Tv.Onelike.NatGeoWild.Stream())
	fmt.Println(config)
}
