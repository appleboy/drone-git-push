package repo

import (
	"os/exec"
	"regexp"
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
func RemotePush(remote, branch string, force, followtags bool) *exec.Cmd {
	return RemotePushNamedBranch(remote, "HEAD", branch, force, followtags)
}

func RemotePullRebaseNamedBranch(remote, branch string) *exec.Cmd {
	cmd := exec.Command(
		"git",
		"pull",
		"--rebase",
		remote,
		branch)

	return cmd
}

var validBranchName = regexp.MustCompile(`^[\w\.\-\/]+$`)

func sanitizeInput(input string) string {
	if isValidInput(input) {
		return input
	}
	return ""
}

func isValidInput(input string) bool {
	return validBranchName.MatchString(input)
}

// RemotePushNamedBranch puchs changes from a local to a remote branch.
func RemotePushNamedBranch(remote, localbranch, branch string, force, followtags bool) *exec.Cmd {
	sanitizedRemote := sanitizeInput(remote)
	sanitizedLocalBranch := sanitizeInput(localbranch)
	sanitizedBranch := sanitizeInput(branch)

	cmd := exec.Command(
		"git",
		"push",
		sanitizedRemote,
		sanitizedLocalBranch+":"+sanitizedBranch,
	)

	if force {
		cmd.Args = append(
			cmd.Args,
			"--force")
	}

	if followtags {
		cmd.Args = append(
			cmd.Args,
			"--follow-tags")
	}

	return cmd
}
