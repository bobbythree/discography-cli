// Package tui: terminal_ui.go
package tui

import (
	"bufio"
	"fmt"
	"os"
)

func Run() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Welcome to Discography-cli. Type an artist or album name to get started.")
		fmt.Print("search> ")

		scanner.Scan()
		input := scanner.Text()
		fmt.Printf("You searched for %s\n", input)

	}
}
