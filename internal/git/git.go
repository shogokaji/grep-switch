package git

import (
	"fmt"
	"os/exec"
	"strings"
)

func GetBranches(keyword string) ([]string, error) {
	cmd := exec.Command("git", "branch", "--list")
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("executing git command failed: %w", err)
	}

	var branches []string
	for _, line := range strings.Split(strings.TrimSpace(string(output)), "\n") {
		branch := strings.TrimSpace(line)
		if branch == "" {
			continue
		}
		if keyword == "" || strings.Contains(branch, keyword) {
			branches = append(branches, strings.TrimPrefix(branch, "* "))
		}
	}

	if len(branches) == 0 && keyword != "" {
		return nil, fmt.Errorf("no branches found matching keyword: %s", keyword)
	}

	return branches, nil
}

func SwitchBranch(branch string) error {
	cmd := exec.Command("git", "switch", branch)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("switching to branch %s failed: %w\nCommand output: %s", branch, err, output)
	}
	return nil
}
