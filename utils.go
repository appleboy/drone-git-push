package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func execute(cmd *exec.Cmd) error {
	fmt.Println("+", strings.Join(cmd.Args, " "))

	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	return cmd.Run()
}
