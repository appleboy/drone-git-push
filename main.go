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
	buildCommit string
)

const (
	DefaultRemoteName = "deploy"
	DefaultLocalRef   = "HEAD"
)

func main() {
	fmt.Printf("Drone Git Push Plugin built from %s\n", buildCommit)

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

	if err := repo.WriteNetrc(workspace); err != nil {
		return err
	}

	if vargs.RemoteName == "" {
		vargs.RemoteName = DefaultRemoteName
	}

	if vargs.LocalBranch == "" {
		vargs.LocalBranch = DefaultLocalRef
	}

	if vargs.Remote != "" {
		cmd := repo.RemoteAdd(
			vargs.RemoteName,
			vargs.Remote)

		if err := execute(cmd, workspace); err != nil {
			return err
		}

		defer func() {
			execute(
				repo.RemoteRemove(
					vargs.RemoteName),
				workspace)
		}()
	}

	if vargs.Commit {
		cmd := repo.ForceAdd()
		if err := execute(cmd, workspace); err != nil {
			return err
		}

		cmd = repo.ForceCommit()
		if err := execute(cmd, workspace); err != nil {
			return err
		}
	}

	cmd := repo.RemotePushNamedBranch(
		vargs.RemoteName,
		vargs.LocalBranch,
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
