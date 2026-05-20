// Package tui: terminal_ui.go
package tui

import (
	"bufio"
	"fmt"
	"os"

	"github.com/bobbythree/discography-cli/data"
	"github.com/bobbythree/discography-cli/search"
)

func Run() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Welcome to Discography-cli. Type an artist or album name to get started.")

	for {
		fmt.Print("search> ")

		scanner.Scan()
		input := scanner.Text()

		// checks
		if len(input) == 0 {
			continue
		}

		result := search.QueryAlbums(input, data.Albums)
		if len(result) == 0 {
			fmt.Println("No albums found. Try again.")
		}

		for _, a := range result {
			fmt.Printf("%s - %s\n", a.Artist, a.AlbumName)
		}
		fmt.Println("")
	}
}
