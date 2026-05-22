// Package tui: terminal_ui.go
package tui

import (
	"bufio"
	"fmt"
	"os"
)

func Run() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Welcome to Discography-cli. Type an artist or album name to get started.")

	for {
		fmt.Print("search> ")

		scanner.Scan()
		input := scanner.Text()

	}
}
