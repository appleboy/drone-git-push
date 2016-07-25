package repo

import (
	"os/exec"
)

// GlobalUser sets the global git author email.
func GlobalUser(email string) *exec.Cmd {
	cmd := exec.Command(
		"git",
		"config",
		"--global",
		"user.email",
		email)

	return cmd
}

// GlobalName sets the global git author name.
func GlobalName(author string) *exec.Cmd {
	cmd := exec.Command(
		"git",
		"config",
		"--global",
		"user.name",
		author)

	return cmd
}
