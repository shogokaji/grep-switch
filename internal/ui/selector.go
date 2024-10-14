package ui

import (
	"fmt"
	"os"

	"github.com/eiannone/keyboard"
	"github.com/shogokaji/grep-switch/internal/util"
)

func Selector(branches []string) (string, error) {
	if len(branches) == 0 {
		return "", fmt.Errorf("no branches to select from")
	} else if len(branches) == 1 {
		return branches[0], nil
	}

	choice := 0

	if err := keyboard.Open(); err != nil {
		return "", fmt.Errorf("failed to open keyboard: %w", err)
	}
	defer cleanUp()

	for {
		displayBranches(branches, choice)

		char, key, err := keyboard.GetKey()
		if err != nil {
			return "", fmt.Errorf("failed to get key: %w", err)
		}

		switch key {
		case keyboard.KeyArrowDown, keyboard.KeyArrowUp:
			choice = move(choice, len(branches), key == keyboard.KeyArrowDown)
		case keyboard.KeyEnter:
			return branches[choice], nil
		case keyboard.KeyEsc:
			return "", fmt.Errorf("selection cancelled by user")
		default:
			switch char {
			case 'j':
				choice = move(choice, len(branches), true)
			case 'k':
				choice = move(choice, len(branches), false)
			case 'q':
				cleanUp()
				os.Exit(0)
			}
		}
	}
}

func displayBranches(branches []string, choice int) {
	util.ClearScreen()
	fmt.Println("Select a branch to switch:")
	fmt.Println("Use arrow keys or j/k to navigate, Enter to select, q to quit")

	for i, branch := range branches {
		if i == choice {
			fmt.Print("\033[1;31m>\033[0m \033[1;32m")
		} else {
			fmt.Print("  ")
		}
		fmt.Printf("%s\033[0m\n", branch)
	}
}

func move(current, max int, down bool) int {
	if down {
		if current < max-1 {
			return current + 1
		}
	} else {
		if current > 0 {
			return current - 1
		}
	}
	return current
}

func cleanUp() {
	util.ClearScreen()
	if err := keyboard.Close(); err != nil {
		fmt.Printf("Warning: failed to close keyboard: %v\n", err)
	}
}
