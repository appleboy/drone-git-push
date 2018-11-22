package repo

import (
	"os/exec"
)

const defaultCommitMessage = "[skip ci] Commit dirty state"

// ForceAdd forces the addition of all dirty files.
func ForceAdd() *exec.Cmd {
	cmd := exec.Command(
		"git",
		"add",
		"--all",
		"--force")

	return cmd
}

// Add updates the index to match the working tree.
func Add() *exec.Cmd {
	cmd := exec.Command(
		"git",
		"add",
		"--all")

	return cmd
}

// TestCleanTree returns non-zero if diff between index and local repository
func TestCleanTree() *exec.Cmd {
	cmd := exec.Command(
		"git",
		"diff-index",
		"--quiet",
		"HEAD",
		"--ignore-submodules")

	return cmd
}

// EmptyCommit simply create an empty commit
func EmptyCommit(msg string, noVerify bool) *exec.Cmd {
	if msg == "" {
		msg = defaultCommitMessage
	}

	cmd := exec.Command(
		"git",
		"commit",
		"--allow-empty",
		"-m",
		msg,
	)

	if noVerify {
		cmd.Args = append(
			cmd.Args,
			"--no-verify")
	}

	return cmd
}

// ForceCommit commits every change while skipping CI.
func ForceCommit(msg string, noVerify bool) *exec.Cmd {
	if msg == "" {
		msg = defaultCommitMessage
	}

	cmd := exec.Command(
		"git",
		"commit",
		"-m",
		msg,
	)

	if noVerify {
		cmd.Args = append(
			cmd.Args,
			"--no-verify")
	}

	return cmd
}
