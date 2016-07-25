package repo

import (
	"os/exec"
)

// ForceAdd forces the addition of all dirty files.
func ForceAdd() *exec.Cmd {
	cmd := exec.Command(
		"git",
		"add",
		"--all",
		"--force")

	return cmd
}

// ForceCommit commits every change while skipping CI.
func ForceCommit() *exec.Cmd {
	cmd := exec.Command(
		"git",
		"commit",
		"-m '[skip ci] Commit dirty state'")

	return cmd
}
