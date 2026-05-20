package main

import (
	"fmt"

	"github.com/bobbythree/discography-cli/data"
)

func main() {
	for _, album := range data.Albums {
		fmt.Printf("%v: %v\n", album.Artist, album.AlbumName)
	}
}
