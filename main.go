package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/appleboy/com/random"
	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
)

// Version set at compile-time
var (
	Version string
)

func main() {
	// Load env-file if it exists first
	if filename, found := os.LookupEnv("PLUGIN_ENV_FILE"); found {
		_ = godotenv.Load(filename)
	}

	if _, err := os.Stat("/run/drone/env"); err == nil {
		_ = godotenv.Overload("/run/drone/env")
	}

	app := cli.NewApp()
	app.Name = "git-push plugin"
	app.Usage = "git-push plugin"
	app.Copyright = "Copyright (c) " + strconv.Itoa(time.Now().Year()) + " Bo-Yi Wu"
	app.Action = run
	app.Version = Version
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "commit.author.name",
			Usage:   "git author name",
			EnvVars: []string{"PLUGIN_AUTHOR_NAME", "DRONE_COMMIT_AUTHOR", "CI_COMMIT_AUTHOR", "INPUT_AUTHOR_NAME"},
		},
		&cli.StringFlag{
			Name:    "commit.author.email",
			Usage:   "git author email",
			EnvVars: []string{"PLUGIN_AUTHOR_EMAIL", "DRONE_COMMIT_AUTHOR_EMAIL", "CI_COMMIT_AUTHOR_EMAIL", "INPUT_AUTHOR_EMAIL"},
		},

		&cli.StringFlag{
			Name:    "netrc.machine",
			Usage:   "netrc machine",
			EnvVars: []string{"PLUGIN_NETRC_MACHINE", "DRONE_NETRC_MACHINE", "INPUT_NETRC_MACHINE"},
		},
		&cli.StringFlag{
			Name:    "netrc.username",
			Usage:   "netrc username",
			EnvVars: []string{"PLUGIN_USERNAME", "DRONE_NETRC_USERNAME", "GITHUB_USERNAME", "INPUT_USERNAME"},
		},
		&cli.StringFlag{
			Name:    "netrc.password",
			Usage:   "netrc password",
			EnvVars: []string{"PLUGIN_PASSWORD", "DRONE_NETRC_PASSWORD", "GITHUB_PASSWORD", "INPUT_PASSWORD"},
		},
		&cli.StringFlag{
			Name:    "ssh-key",
			Usage:   "private ssh key",
			EnvVars: []string{"PLUGIN_SSH_KEY", "GIT_PUSH_SSH_KEY", "INPUT_SSH_KEY"},
		},
		&cli.StringFlag{
			Name:    "remote",
			Usage:   "url of the remote repo",
			EnvVars: []string{"PLUGIN_REMOTE", "GIT_PUSH_REMOTE", "INPUT_REMOTE"},
		},
		&cli.StringFlag{
			Name:    "remote-name",
			Usage:   "name of the remote repo",
			Value:   "",
			EnvVars: []string{"PLUGIN_REMOTE_NAME", "GIT_PUSH_REMOTE_NAME", "INPUT_REMOTE_NAME"},
		},
		&cli.StringFlag{
			Name:    "branch",
			Usage:   "name of remote branch",
			EnvVars: []string{"PLUGIN_BRANCH", "GIT_PUSH_BRANCH", "INPUT_BRANCH"},
			Value:   "master",
		},
		&cli.StringFlag{
			Name:    "local-branch",
			Usage:   "name of local branch",
			Value:   "HEAD",
			EnvVars: []string{"PLUGIN_LOCAL_BRANCH", "GIT_PUSH_LOCAL_BRANCH", "INPUT_LOCAL_BRANCH"},
		},
		&cli.StringFlag{
			Name:    "path",
			Usage:   "path to git repo",
			EnvVars: []string{"PLUGIN_PATH", "GIT_PUSH_PATH", "INPUT_PATH"},
		},
		&cli.BoolFlag{
			Name:    "force",
			Usage:   "force push to remote",
			EnvVars: []string{"PLUGIN_FORCE", "GIT_PUSH_FORCE", "INPUT_FORCE"},
		},
		&cli.BoolFlag{
			Name:    "followtags",
			Usage:   "push to remote with tags",
			EnvVars: []string{"PLUGIN_FOLLOWTAGS", "GIT_PUSH_FOLLOWTAGS", "INPUT_FOLLOWTAGS"},
		},
		&cli.BoolFlag{
			Name:    "skip-verify",
			Usage:   "skip ssl verification",
			EnvVars: []string{"PLUGIN_SKIP_VERIFY", "GIT_PUSH_SKIP_VERIFY", "INPUT_SKIP_VERIFY"},
		},
		&cli.BoolFlag{
			Name:    "commit",
			Usage:   "commit dirty changes",
			EnvVars: []string{"PLUGIN_COMMIT", "GIT_PUSH_COMMIT", "INPUT_COMMIT"},
		},
		&cli.StringFlag{
			Name:    "commit-message",
			Usage:   "commit message",
			EnvVars: []string{"PLUGIN_COMMIT_MESSAGE", "GIT_PUSH_COMMIT_MESSAGE", "INPUT_COMMIT_MESSAGE"},
		},
		&cli.StringFlag{
			Name:    "tag",
			Usage:   "tag to add to the commit",
			EnvVars: []string{"PLUGIN_TAG", "GIT_PUSH_TAG", "INPUT_TAG"},
		},
		&cli.BoolFlag{
			Name:    "empty-commit",
			Usage:   "empty commit",
			EnvVars: []string{"PLUGIN_EMPTY_COMMIT", "GIT_PUSH_EMPTY_COMMIT", "INPUT_EMPTY_COMMIT"},
		},
		&cli.BoolFlag{
			Name:    "no-verify",
			Usage:   "bypasses the pre-commit and commit-msg hooks",
			EnvVars: []string{"PLUGIN_NO_VERIFY", "GIT_PUSH_NO_VERIFY", "INPUT_NO_VERIFY"},
		},
		&cli.BoolFlag{
			Name:    "rebase",
			Usage:   "pull rebase branch before push",
			EnvVars: []string{"PLUGIN_REBASE", "GIT_PUSH_REBASE", "INPUT_REBASE"},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(c *cli.Context) error {
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
			Key:           c.String("ssh-key"),
			Remote:        c.String("remote"),
			RemoteName:    c.String("remote-name"),
			Branch:        c.String("branch"),
			LocalBranch:   c.String("local-branch"),
			Path:          c.String("path"),
			Force:         c.Bool("force"),
			FollowTags:    c.Bool("followtags"),
			SkipVerify:    c.Bool("skip-verify"),
			Commit:        c.Bool("commit"),
			CommitMessage: c.String("commit-message"),
			Tag:           c.String("tag"),
			EmptyCommit:   c.Bool("empty-commit"),
			NoVerify:      c.Bool("no-verify"),
			Rebase:        c.Bool("rebase"),
		},
	}

	if plugin.Config.RemoteName == "" {
		plugin.Config.RemoteName = random.String(10)
	}

	return plugin.Exec()
}
