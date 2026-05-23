// Package tui: terminal_ui.go
package tui

import (
	"bufio"
	"fmt"
	"os"

	"github.com/bobbythree/discography-cli/discogsapi"
)

func Run() {
	scanner := bufio.NewScanner(os.Stdin)

	client := discogsapi.NewClient()

	fmt.Println("Welcome to Discography-cli. Type an album name to get started.")

	for {
		fmt.Print("search> ")

		scanner.Scan()

		input := scanner.Text()

		results, err := client.AlbumSearch(input)
		if err != nil {
			fmt.Println("error: ", err)
			continue
		}

		if len(results) == 0 {
			fmt.Println("No results found.")
			continue
		}

		for i, album := range results {
			fmt.Printf("%d. %s (%s)\n", i+1, album.Title, album.Year)
		}

		fmt.Println()
	}
}
