package repo

import (
	"os/exec"
)

// SkipVerify disables globally the git ssl verification.
func SkipVerify() *exec.Cmd {
	cmd := exec.Command(
		"git",
		"config",
		"--global",
		"http.sslVerify",
		"false")

	return cmd
}
