package repo

import (
	"os/exec"
)

func RemoteRemove(name string) *exec.Cmd {
	cmd := exec.Command(
		"git",
		"remote",
		"rm",
		name)

	return cmd
}

func RemoteAdd(name, url string) *exec.Cmd {
	cmd := exec.Command(
		"git",
		"remote",
		"add",
		name,
		url)

	return cmd
}

func RemotePush(remote, branch string, force bool) *exec.Cmd {
	return RemotePushNamedBranch(remote, "HEAD", branch, force)
}

func RemotePushNamedBranch(remote, localbranch string, branch string, force bool) *exec.Cmd {
	cmd := exec.Command(
		"git",
		"push",
		remote,
		localbranch+":"+branch)

	if force {
		cmd.Args = append(
			cmd.Args,
			"--force")
	}

	return cmd
}
