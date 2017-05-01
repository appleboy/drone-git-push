package repo

import (
	"fmt"
	"os/exec"
)

const defaultCommitMessage = "[skip ci] Commit dirty state"

// ForceAdd forces the addition of all dirty files.
func ForceAdd() *exec.Cmd {
	cmd := exec.Command(
		"git",
		"add",
		"--all",
		"--force")

	return cmd
}

// EmptyCommit simply create an empty commit
func EmptyCommit(msg string) *exec.Cmd {
	commitMessage := defaultCommitMessage

	if msg != "" {
		commitMessage = msg
	}

	cmd := exec.Command(
		"git",
		"commit",
		"--allow-empty",
		fmt.Sprintf("-m '%s'", commitMessage))

	return cmd
}

// ForceCommit commits every change while skipping CI.
func ForceCommit(msg string) *exec.Cmd {
	commitMessage := defaultCommitMessage

	if msg != "" {
		commitMessage = msg
	}

	cmd := exec.Command(
		"git",
		"commit",
		fmt.Sprintf("-m '%s'", commitMessage))

	return cmd
}
