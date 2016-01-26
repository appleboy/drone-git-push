package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/drone-plugins/drone-git-push/repo"
	"github.com/drone/drone-go/drone"
	"github.com/drone/drone-go/plugin"
)

var (
	buildDate string
)

func main() {
	fmt.Printf("Drone Git Push Plugin built at %s\n", buildDate)

	workspace := drone.Workspace{}
	build := drone.Build{}
	vargs := Params{}

	plugin.Param("workspace", &workspace)
	plugin.Param("build", &build)
	plugin.Param("vargs", &vargs)
	plugin.MustParse()

	err := run(&workspace, &build, &vargs)

	if err != nil {
		fmt.Println(err)

		os.Exit(1)
		return
	}
}

func run(workspace *drone.Workspace, build *drone.Build, vargs *Params) error {
	repo.GlobalName(build).Run()
	repo.GlobalUser(build).Run()

	if vargs.SkipVerify {
		repo.SkipVerify().Run()
	}

	if err := repo.WriteKey(workspace); err != nil {
		return err
	}

	defer func() {
		execute(
			repo.RemoteRemove(
				"deploy"),
			workspace)
	}()

	cmd := repo.RemoteAdd(
		"deploy",
		vargs.Remote)

	if err := execute(cmd, workspace); err != nil {
		return err
	}

	cmd = repo.RemotePush(
		"deploy",
		vargs.Branch,
		vargs.Force)

	if err := execute(cmd, workspace); err != nil {
		return err
	}

	return nil
}

func execute(cmd *exec.Cmd, workspace *drone.Workspace) error {
	trace(cmd)

	cmd.Dir = workspace.Path
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	return cmd.Run()
}

func trace(cmd *exec.Cmd) {
	fmt.Println("$", strings.Join(cmd.Args, " "))
}
