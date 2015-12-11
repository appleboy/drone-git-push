package repo

import (
	"os/exec"
)

func ForceAdd() *exec.Cmd {
	cmd := exec.Command(
		"git",
		"add",
		"--all",
		"--force")

	return cmd
}

func ForceCommit() *exec.Cmd {
	cmd := exec.Command(
		"git",
		"commit",
		"-m 'Commit dirty state'")

	return cmd
}
