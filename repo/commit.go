package repo

import (
	"context"
	"fmt"
	"os/exec"
)

const defaultCommitMessage = "[skip ci] Commit dirty state"

// ForceAdd forces the addition of all dirty files.
func ForceAdd(ctx context.Context) *exec.Cmd {
	cmd := exec.CommandContext(
		ctx,
		"git",
		"add",
		"--all",
		"--force")

	return cmd
}

// Add updates the index to match the working tree.
func Add(ctx context.Context) *exec.Cmd {
	cmd := exec.CommandContext(
		ctx,
		"git",
		"add",
		"--all")

	return cmd
}

// Tag add tag to the working tree.
func Tag(ctx context.Context, tag string) *exec.Cmd {
	cmd := exec.CommandContext(
		ctx,
		"git",
		"tag",
		"-a",
		tag,
		"-m",
		tag)

	return cmd
}

// TestCleanTree returns non-zero if diff between index and local repository
func TestCleanTree(ctx context.Context) *exec.Cmd {
	cmd := exec.CommandContext(
		ctx,
		"git",
		"diff-index",
		"--quiet",
		"HEAD",
		"--ignore-submodules")

	return cmd
}

// EmptyCommit simply create an empty commit
func EmptyCommit(
	ctx context.Context,
	msg string,
	noVerify bool,
	authorName, authorEmail string,
) *exec.Cmd {
	if msg == "" {
		msg = defaultCommitMessage
	}

	cmd := exec.CommandContext(
		ctx,
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

	if authorName != "" || authorEmail != "" {
		cmd.Args = append(
			cmd.Args,
			fmt.Sprintf("--author=\"%s <%q>\"", authorName, authorEmail))
	}

	return cmd
}

// ForceCommit commits every change while skipping CI.
func ForceCommit(
	ctx context.Context,
	msg string,
	noVerify bool,
	authorName, authorEmail string,
) *exec.Cmd {
	if msg == "" {
		msg = defaultCommitMessage
	}

	cmd := exec.CommandContext(
		ctx,
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

	if authorName != "" || authorEmail != "" {
		cmd.Args = append(
			cmd.Args,
			fmt.Sprintf("--author=\"%s <%s>\"", authorName, authorEmail))
	}

	return cmd
}
