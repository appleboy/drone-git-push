package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func execute(cmd *exec.Cmd) error {
	//nolint:forbidigo // Preserve the plugin CLI output format on stdout.
	fmt.Println(
		"+",
		strings.Join(cmd.Args, " "),
	)

	cmd.Env = os.Environ()
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	return cmd.Run()
}
