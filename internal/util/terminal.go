package util

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func DisplayErrorBox(message string) {
	width := len(message) + 4
	horizontalBorder := strings.Repeat("─", width)
	
	fmt.Printf("\033[1;31m")
	fmt.Printf("┌%s┐\n", horizontalBorder)
	fmt.Printf("│  %s  │\n", message)
	fmt.Printf("└%s┘\n", horizontalBorder)
	fmt.Printf("\033[0m")
}

func ClearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}