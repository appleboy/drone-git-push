package repo

import (
	"context"
	"os/exec"
)

// GlobalUser sets the global git author email.
func GlobalUser(ctx context.Context, email string) *exec.Cmd {
	cmd := exec.CommandContext(
		ctx,
		"git",
		"config",
		// "--global",
		"user.email",
		email)

	return cmd
}

// GlobalName sets the global git author name.
func GlobalName(ctx context.Context, author string) *exec.Cmd {
	cmd := exec.CommandContext(
		ctx,
		"git",
		"config",
		// "--global",
		"user.name",
		author)

	return cmd
}
