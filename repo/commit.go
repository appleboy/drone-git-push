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
		"-m '[skip ci] Commit dirty state'") // skip the CI build since this commit was triggered by the build system, not by a user

	return cmd
}
