package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// helper function to run git commands
func runGitCommand(args ...string) {
	cmd := exec.Command("git", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println("Error:", err)
	}
}

func main() {
	var rootCmd = &cobra.Command{
		Use:   "gitundo",
		Short: "GitUndo - simplify git history operations",
	}

	// Command: undo last commit
	var undoLastCmd = &cobra.Command{
		Use:   "last-commit",
		Short: "Undo the last commit",
		Run: func(cmd *cobra.Command, args []string) {
			runGitCommand("reset", "--hard", "HEAD~1")
		},
	}

	rootCmd.AddCommand(undoLastCmd)
	rootCmd.Execute()
}
