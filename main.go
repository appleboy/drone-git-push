package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/drone/drone-go/drone"
	"github.com/drone/drone-go/plugin"
)

// Params stores the git push parameters used to
// configure and customzie the git push behavior.
type Params struct {
	Remote     string `json:"remote"`
	Branch     string `json:"branch"`
	Force      bool   `json:"force"`
	SkipVerify bool   `json:"skip_verify"`
}

func main() {
	v := new(Params)
	b := new(drone.Build)
	w := new(drone.Workspace)
	plugin.Param("build", b)
	plugin.Param("workspace", w)
	plugin.Param("vargs", &v)
	plugin.MustParse()

	err := run(b, w, v)
	if err != nil {
		os.Exit(1)
	}
}

func run(b *drone.Build, w *drone.Workspace, v *Params) error {

	// write the rsa private key if provided
	if err := writeKey(w); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return err
	}

	// sets the global user and name
	globalName(b).Run()
	globalUser(b).Run()

	defer func() {
		// removes the remote to avoid conflict
		cmd := remoteRemove("deploy")
		cmd.Dir = w.Path
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin
		cmd.Run()
	}()

	// if the user is pushing to a self-hosted service
	// with a self-signed certificate this may be required.
	if v.SkipVerify {
		skipVerify().Run()
	}

	// adds the remote
	cmd := remoteAdd("deploy", v.Remote)
	cmd.Dir = w.Path
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	trace(cmd)
	cmd.Run()

	// pushes our code
	cmd = push("deploy", v.Branch, v.Force)
	cmd.Dir = w.Path
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	trace(cmd)
	return cmd.Run()
}

func remoteAdd(name, url string) *exec.Cmd {
	return exec.Command(
		"git",
		"remote",
		"add",
		name,
		url,
	)
}

func remoteRemove(name string) *exec.Cmd {
	return exec.Command(
		"git",
		"remote",
		"rm",
		name,
	)
}

func globalUser(b *drone.Build) *exec.Cmd {
	return exec.Command(
		"git",
		"config",
		"--global",
		"user.email",
		b.Email,
	)
}

func globalName(b *drone.Build) *exec.Cmd {
	return exec.Command(
		"git",
		"config",
		"--global",
		"user.name",
		b.Author,
	)
}

func push(remote, branch string, force bool) *exec.Cmd {
	cmd := exec.Command(
		"git",
		"push",
		"remote",
		"HEAD:"+branch,
	)
	if force {
		cmd.Args = append(cmd.Args, "--force")
	}
	return cmd
}

// skipVerify returns a git command that, when executed
// configures git to skip ssl verification. This should
// may be used with self-signed certificates.
func skipVerify() *exec.Cmd {
	return exec.Command(
		"git",
		"config",
		"--global",
		"http.sslVerify",
		"false",
	)
}

// Trace writes each command to standard error (preceded by a ‘$ ’) before it
// is executed. Used for debugging your build.
func trace(cmd *exec.Cmd) {
	fmt.Println("$", strings.Join(cmd.Args, " "))
}

// Writes the RSA private key
func writeKey(in *drone.Workspace) error {
	if in.Keys == nil || len(in.Keys.Private) == 0 {
		return nil
	}
	home := "/root"
	u, err := user.Current()
	if err == nil {
		home = u.HomeDir
	}
	sshpath := filepath.Join(home, ".ssh")
	if err := os.MkdirAll(sshpath, 0700); err != nil {
		return err
	}
	confpath := filepath.Join(sshpath, "config")
	privpath := filepath.Join(sshpath, "id_rsa")
	ioutil.WriteFile(confpath, []byte("StrictHostKeyChecking no\n"), 0700)
	return ioutil.WriteFile(privpath, []byte(in.Keys.Private), 0600)
}
