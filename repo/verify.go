package repo

import (
	"context"
	"os/exec"
)

// SkipVerify disables globally the git ssl verification.
func SkipVerify(ctx context.Context) *exec.Cmd {
	cmd := exec.CommandContext(
		ctx,
		"git",
		"config",
		"--global",
		"http.sslVerify",
		"false")

	return cmd
}
