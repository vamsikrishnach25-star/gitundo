package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// helper to run git commands
func runGitCommand(args ...string) {
	cmd := exec.Command("git", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		color.Red("Error: %v\n", err)
	}
}

// helper to confirm destructive actions
func confirmAction(prompt string) bool {
	fmt.Printf("%s (y/N): ", prompt)
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input := strings.ToLower(strings.TrimSpace(scanner.Text()))
		return input == "y" || input == "yes"
	}
	return false
}

func main() {
	var rootCmd = &cobra.Command{
		Use:   "gitundo",
		Short: "GitUndo - simplify git history operations",
	}

	// Undo last commit
	var undoLastCmd = &cobra.Command{
		Use:   "last-commit",
		Short: "Undo the last commit",
		Run: func(cmd *cobra.Command, args []string) {
			if confirmAction("This will delete the last commit! Are you sure?") {
				runGitCommand("reset", "--hard", "HEAD~1")
			} else {
				color.Yellow("Aborted.")
			}
		},
	}

	// Undo file changes
	var undoFileCmd = &cobra.Command{
		Use:   "file [filename]",
		Short: "Undo changes to a file",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			runGitCommand("checkout", "--", args[0])
		},
	}

	// Undo last staged file
	var undoStagedCmd = &cobra.Command{
		Use:   "unstage [filename]",
		Short: "Unstage a file",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			runGitCommand("reset", "HEAD", args[0])
		},
	}

	// Revert a specific commit
	var revertCmd = &cobra.Command{
		Use:   "revert [commit-hash]",
		Short: "Revert a specific commit",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if confirmAction("This will create a new commit reverting the changes. Proceed?") {
				runGitCommand("git", "revert", args[0])
			} else {
				color.Yellow("Aborted.")
			}
		},
	}

	// Show recent commits
	var showCommitsCmd = &cobra.Command{
		Use:   "log",
		Short: "Show recent commits",
		Run: func(cmd *cobra.Command, args []string) {
			runGitCommand("log", "--oneline", "-n", "10")
		},
	}

	// Register commands
	rootCmd.AddCommand(undoLastCmd, undoFileCmd, undoStagedCmd, revertCmd, showCommitsCmd)

	rootCmd.Execute()
}
