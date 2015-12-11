package repo

import (
	"os/exec"

	"github.com/drone/drone-go/drone"
)

func GlobalUser(build *drone.Build) *exec.Cmd {
	cmd := exec.Command(
		"git",
		"config",
		"--global",
		"user.email",
		build.Email)

	return cmd
}

func GlobalName(build *drone.Build) *exec.Cmd {
	cmd := exec.Command(
		"git",
		"config",
		"--global",
		"user.name",
		build.Author)

	return cmd
}
