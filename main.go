package main

import (
	"fmt"
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/joho/godotenv"
	"github.com/urfave/cli"
)

var build = "0" // build number set at compile-time

func main() {
	app := cli.NewApp()
	app.Name = "git-push plugin"
	app.Usage = "git-push plugin"
	app.Action = run
	app.Version = fmt.Sprintf("1.0.%s", build)
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "commit.author.name",
			Usage:  "git author name",
			EnvVar: "DRONE_COMMIT_AUTHOR",
		},
		cli.StringFlag{
			Name:   "commit.author.email",
			Usage:  "git author email",
			EnvVar: "DRONE_COMMIT_AUTHOR_EMAIL",
		},

		cli.StringFlag{
			Name:   "netrc.machine",
			Usage:  "netrc machine",
			EnvVar: "DRONE_NETRC_MACHINE",
		},
		cli.StringFlag{
			Name:   "netrc.username",
			Usage:  "netrc username",
			EnvVar: "DRONE_NETRC_USERNAME",
		},
		cli.StringFlag{
			Name:   "netrc.password",
			Usage:  "netrc password",
			EnvVar: "DRONE_NETRC_PASSWORD",
		},
		cli.StringFlag{
			Name:   "ssh-key",
			Usage:  "private ssh key",
			EnvVar: "PLUGIN_SSH_KEY,GIT_PUSH_SSH_KEY",
		},
		cli.StringFlag{
			Name:   "remote",
			Usage:  "url of the remote repo",
			EnvVar: "PLUGIN_REMOTE,GIT_PUSH_REMOTE",
		},
		cli.StringFlag{
			Name:   "remote-name",
			Usage:  "name of the remote repo",
			Value:  "deploy",
			EnvVar: "PLUGIN_REMOTE_NAME,GIT_PUSH_REMOTE_NAME",
		},
		cli.StringFlag{
			Name:   "branch",
			Usage:  "name of remote branch",
			EnvVar: "PLUGIN_BRANCH,GIT_PUSH_BRANCH",
		},
		cli.StringFlag{
			Name:   "local-branch",
			Usage:  "name of local branch",
			Value:  "HEAD",
			EnvVar: "PLUGIN_LOCAL_BRANCH,GIT_PUSH_LOCAL_BRANCH",
		},
		cli.BoolFlag{
			Name:   "force",
			Usage:  "force push to remote",
			EnvVar: "PLUGIN_FORCE,GIT_PUSH_FORCE",
		},
		cli.BoolFlag{
			Name:   "skip-verify",
			Usage:  "skip ssl verification",
			EnvVar: "PLUGIN_SKIP_VERIFY,GIT_PUSH_SKIP_VERIFY",
		},
		cli.BoolFlag{
			Name:   "commit",
			Usage:  "commit dirty changes",
			EnvVar: "PLUGIN_COMMIT,GIT_PUSH_COMMIT",
		},
		cli.StringFlag{
			Name:  "env-file",
			Usage: "source env file",
		},
	}

	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}

func run(c *cli.Context) error {
	if c.String("env-file") != "" {
		_ = godotenv.Load(c.String("env-file"))
	}

	plugin := Plugin{
		Netrc: Netrc{
			Login:    c.String("netrc.username"),
			Machine:  c.String("netrc.machine"),
			Password: c.String("netrc.password"),
		},
		Commit: Commit{
			Author: Author{
				Name:  c.String("commit.author.name"),
				Email: c.String("commit.author.email"),
			},
		},
		Config: Config{
			Key:         c.String("ssh-key"),
			Remote:      c.String("remote"),
			RemoteName:  c.String("remote-name"),
			Branch:      c.String("branch"),
			LocalBranch: c.String("local-branch"),
			Force:       c.Bool("force"),
			SkipVerify:  c.Bool("skip-verify"),
			Commit:      c.Bool("commit"),
		},
	}

	return plugin.Exec()
}
