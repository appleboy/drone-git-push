package repo

import (
	"os/exec"
)

func SkipVerify() *exec.Cmd {
	cmd := exec.Command(
		"git",
		"config",
		"--global",
		"http.sslVerify",
		"false")

	return cmd
}
