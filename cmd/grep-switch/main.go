package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/shogokaji/grep-switch/internal/git"
	"github.com/shogokaji/grep-switch/internal/ui"
	"github.com/shogokaji/grep-switch/internal/util"
)

func main() {
	for {
		if err := run(); err != nil {
			util.DisplayErrorBox(err.Error())
			if !shouldRetry() {
				break
			}
		} else {
			break
		}
	}
}

func run() error {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter a search term: ")
	keyword, err := reader.ReadString('\n')
	if err != nil {
		return fmt.Errorf("failed to read input: %w", err)
	}
	keyword = strings.TrimSpace(keyword)

	branches, err := git.GetBranches(keyword)
	if err != nil {
		return fmt.Errorf("failed to get branches: %w", err)
	}

	if len(branches) == 0 {
		return fmt.Errorf("no matching branches found")
	}

	selectedBranch, err := ui.Selector(branches)
	if err != nil {
		return fmt.Errorf("branch selection failed: %w", err)
	}
	if err := git.SwitchBranch(selectedBranch); err != nil {
		return fmt.Errorf("failed to switch branch: %w", err)
	}

	util.ClearScreen()
	fmt.Printf("Switched to branch '%s'\n", selectedBranch)
	return nil
}

func shouldRetry() bool {
	for {
		fmt.Print("Do you want to try again? (y/n): ")
		reader := bufio.NewReader(os.Stdin)
		response, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input. Please try again.")
			continue
		}
		response = strings.TrimSpace(strings.ToLower(response))
		util.ClearScreen()
		switch response {
		case "y", "yes":
			return true
		case "n", "no":
			return false
		default:
			fmt.Println("Invalid input. Please enter 'y' or 'n'.")
		}
	}
}