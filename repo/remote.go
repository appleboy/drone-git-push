package repo

import (
	"os/exec"
)

// RemoteRemove drops the defined remote from a git repo.
func RemoteRemove(name string) *exec.Cmd {
	cmd := exec.Command(
		"git",
		"remote",
		"rm",
		name)

	return cmd
}

// RemoteAdd adds an additional remote to a git repo.
func RemoteAdd(name, url string) *exec.Cmd {
	cmd := exec.Command(
		"git",
		"remote",
		"add",
		name,
		url)

	return cmd
}

// RemotePush pushs the changes from the local head to a remote branch..
func RemotePush(remote, branch string, force bool) *exec.Cmd {
	return RemotePushNamedBranch(remote, "HEAD", branch, force)
}

// RemotePushNamedBranch puchs changes from a local to a remote branch.
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
