package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Workshop chapters order
var order = []string{
	"main",
	"001-initial-setup",
	"002-mcp-weather-definition",
	"002-mcp-weather-complete",
	"003-mcp-zoo",
	"003-mcp-zoo-complete",
}

func main() {
	if len(os.Args) < 2 {
		showHelp()
		return
	}

	command := os.Args[1]

	switch command {
	case "next":
		moveToNext()
	case "status":
		showStatus()
	case "jump":
		if len(os.Args) < 3 {
			fmt.Println("Error: Please specify a chapter name")
			showAvailableChapters()
			os.Exit(1)
		}
		jumpToChapter(os.Args[2])
	case "reset":
		resetToMain()
	case "help":
		showHelp()
	default:
		fmt.Printf("Unknown command: %s\n\n", command)
		showHelp()
		os.Exit(1)
	}
}

func checkGitRepo() {
	cmd := exec.Command("git", "rev-parse", "--git-dir")
	if err := cmd.Run(); err != nil {
		fmt.Println("Error: Not in a Git repository!")
		os.Exit(1)
	}
}

func getCurrentBranch() (string, error) {
	cmd := exec.Command("git", "branch", "--show-current")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(output)), nil
}

func findCurrentPosition(currentBranch string) int {
	for i, branch := range order {
		if branch == currentBranch {
			return i
		}
	}
	return -1
}

func branchExists(branchName string) bool {
	cmd := exec.Command("git", "show-ref", "--verify", "--quiet", "refs/heads/"+branchName)
	return cmd.Run() == nil
}

func runGitCommand(args ...string) error {
	cmd := exec.Command("git", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func moveToNext() {
	checkGitRepo()

	currentBranch, err := getCurrentBranch()
	if err != nil {
		fmt.Printf("Error getting current branch: %v\n", err)
		os.Exit(1)
	}

	currentPos := findCurrentPosition(currentBranch)
	if currentPos == -1 {
		fmt.Printf("Current branch '%s' is not in the workshop order.\n", currentBranch)
		showAvailableChapters()
		os.Exit(1)
	}

	nextPos := currentPos + 1
	if nextPos >= len(order) {
		fmt.Printf("You're already at the last chapter: %s\n", currentBranch)
		fmt.Println("Workshop completed! 🎉")
		return
	}

	nextBranch := order[nextPos]

	fmt.Printf("Moving from chapter: %s\n", currentBranch)
	fmt.Printf("Moving to chapter: %s\n", nextBranch)
	fmt.Println()

	// Discard all changes in current branch
	fmt.Println("Discarding all changes in current branch...")
	if err := runGitCommand("reset", "--hard", "HEAD"); err != nil {
		fmt.Printf("Error resetting changes: %v\n", err)
		os.Exit(1)
	}
	if err := runGitCommand("clean", "-fd"); err != nil {
		fmt.Printf("Error cleaning files: %v\n", err)
		os.Exit(1)
	}

	// Switch to or create the next branch
	if branchExists(nextBranch) {
		fmt.Printf("Switching to existing branch: %s\n", nextBranch)
		if err := runGitCommand("checkout", nextBranch); err != nil {
			fmt.Printf("Error switching to branch: %v\n", err)
			os.Exit(1)
		}
	} else {
		fmt.Printf("Branch '%s' does not exist. Creating it...\n", nextBranch)
		if err := runGitCommand("checkout", "-b", nextBranch); err != nil {
			fmt.Printf("Error creating branch: %v\n", err)
			os.Exit(1)
		}
	}

	fmt.Println()
	fmt.Printf("✅ Successfully moved to chapter: %s\n", nextBranch)
	fmt.Printf("Progress: %d/%d\n", nextPos, len(order)-1)
}

func showStatus() {
	checkGitRepo()

	currentBranch, err := getCurrentBranch()
	if err != nil {
		fmt.Printf("Error getting current branch: %v\n", err)
		os.Exit(1)
	}

	currentPos := findCurrentPosition(currentBranch)

	fmt.Println("🎓 Workshop Status")
	fmt.Println("==================")
	fmt.Printf("Current chapter: %s\n", currentBranch)

	if currentPos != -1 {
		fmt.Printf("Progress: %d/%d\n", currentPos+1, len(order))
		fmt.Println()
		fmt.Println("Chapter order:")
		for i, branch := range order {
			marker := "  "
			if i == currentPos {
				marker = "👉"
			} else if i < currentPos {
				marker = "✅"
			}
			fmt.Printf("%s %s\n", marker, branch)
		}
	} else {
		fmt.Println("⚠️  Current branch is not part of the workshop order")
		fmt.Println()
		showAvailableChapters()
	}
}

func jumpToChapter(targetChapter string) {
	checkGitRepo()

	// Check if target chapter is in order
	found := false
	for _, branch := range order {
		if branch == targetChapter {
			found = true
			break
		}
	}

	if !found {
		fmt.Printf("Error: Chapter '%s' not found in workshop order\n", targetChapter)
		showAvailableChapters()
		os.Exit(1)
	}

	fmt.Printf("Jumping to chapter: %s\n", targetChapter)
	fmt.Println()

	// Discard all changes in current branch
	fmt.Println("Discarding all changes in current branch...")
	if err := runGitCommand("reset", "--hard", "HEAD"); err != nil {
		fmt.Printf("Error resetting changes: %v\n", err)
		os.Exit(1)
	}
	if err := runGitCommand("clean", "-fd"); err != nil {
		fmt.Printf("Error cleaning files: %v\n", err)
		os.Exit(1)
	}

	// Switch to target branch
	if branchExists(targetChapter) {
		fmt.Printf("Switching to existing branch: %s\n", targetChapter)
		if err := runGitCommand("checkout", targetChapter); err != nil {
			fmt.Printf("Error switching to branch: %v\n", err)
			os.Exit(1)
		}
	} else {
		fmt.Printf("Branch '%s' does not exist. Creating it...\n", targetChapter)
		if err := runGitCommand("checkout", "-b", targetChapter); err != nil {
			fmt.Printf("Error creating branch: %v\n", err)
			os.Exit(1)
		}
	}

	fmt.Println()
	fmt.Printf("✅ Successfully jumped to chapter: %s\n", targetChapter)
}

func showAvailableChapters() {
	fmt.Println("Available chapters:")
	for _, branch := range order {
		fmt.Printf("  - %s\n", branch)
	}
}

func showHelp() {
	fmt.Println("🎓 Workshop Navigation Tool")
	fmt.Println("============================")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  workshop next                    - Move to next chapter")
	fmt.Println("  workshop status                  - Show current status")
	fmt.Println("  workshop jump <chapter-name>     - Jump to specific chapter")
	fmt.Println("  workshop reset                   - Reset to main branch (start over)")
	fmt.Println("  workshop help                    - Show this help")
	fmt.Println()
	fmt.Println("Available chapters:")
	for i, branch := range order {
		fmt.Printf("  %d. %s\n", i+1, branch)
	}
	fmt.Println()
	fmt.Println("Note: This tool will discard all uncommitted changes when switching chapters!")
}

func resetToMain() {
	checkGitRepo()

	currentBranch, err := getCurrentBranch()
	if err != nil {
		fmt.Printf("Error getting current branch: %v\n", err)
		os.Exit(1)
	}

	if currentBranch == "main" {
		fmt.Println("You're already on the main branch!")
		return
	}

	fmt.Printf("Resetting workshop: Going back to main branch from %s\n", currentBranch)
	fmt.Println()

	// Discard all changes in current branch
	fmt.Println("Discarding all changes in current branch...")
	if err := runGitCommand("reset", "--hard", "HEAD"); err != nil {
		fmt.Printf("Error resetting changes: %v\n", err)
		os.Exit(1)
	}
	if err := runGitCommand("clean", "-fd"); err != nil {
		fmt.Printf("Error cleaning files: %v\n", err)
		os.Exit(1)
	}

	// Switch back to main branch
	fmt.Println("Switching to main branch...")
	if err := runGitCommand("checkout", "main"); err != nil {
		fmt.Printf("Error switching to main branch: %v\n", err)
		os.Exit(1)
	}

	fmt.Println()
	fmt.Printf("✅ Successfully reset to main branch\n")
	fmt.Println("🎓 Ready to start the workshop from the beginning!")
}
