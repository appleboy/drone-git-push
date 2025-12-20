package main

import (
	"context"
	"os"

	"github.com/appleboy/drone-git-push/repo"
)

type (
	// Netrc structure
	Netrc struct {
		Machine  string
		Login    string
		Password string
	}

	// Commit structure
	Commit struct {
		Author Author
	}

	// Author structure
	Author struct {
		Name  string
		Email string
	}

	// Config structure
	Config struct {
		Key           string
		Remote        string
		RemoteName    string
		Branch        string
		LocalBranch   string
		Path          string
		Force         bool
		FollowTags    bool
		SkipVerify    bool
		Commit        bool
		CommitMessage string
		Tag           string
		EmptyCommit   bool
		NoVerify      bool
		Rebase        bool
	}

	// Plugin Structure
	Plugin struct {
		Netrc  Netrc
		Commit Commit
		Config Config
	}
)

// Exec starts the plugin execution.
func (p Plugin) Exec(ctx context.Context) error {
	if err := p.HandlePath(); err != nil {
		return err
	}

	if err := p.WriteConfig(ctx); err != nil {
		return err
	}

	if err := p.WriteKey(); err != nil {
		return err
	}

	if err := p.WriteNetrc(); err != nil {
		return err
	}

	if err := p.WriteToken(); err != nil {
		return err
	}

	if err := p.HandleCommit(ctx); err != nil {
		return err
	}

	if err := p.HandleTag(ctx); err != nil {
		return err
	}

	if err := p.HandleRemote(ctx); err != nil {
		return err
	}

	if err := p.HandleRebase(ctx); err != nil {
		return err
	}

	if err := p.HandlePush(ctx); err != nil {
		return err
	}

	return p.HandleCleanup(ctx)
}

// WriteConfig writes all required configurations.
func (p Plugin) WriteConfig(ctx context.Context) error {
	if err := execute(repo.GlobalName(ctx, p.Commit.Author.Name)); err != nil {
		return err
	}

	if err := execute(repo.GlobalUser(ctx, p.Commit.Author.Email)); err != nil {
		return err
	}

	if p.Config.SkipVerify {
		if err := execute(repo.SkipVerify(ctx)); err != nil {
			return err
		}
	}

	return nil
}

// WriteKey writes the private SSH key.
func (p Plugin) WriteKey() error {
	return repo.WriteKey(
		p.Config.Key,
	)
}

// WriteNetrc writes the netrc config.
func (p Plugin) WriteNetrc() error {
	return repo.WriteNetrc(
		p.Netrc.Machine,
		p.Netrc.Login,
		p.Netrc.Password,
	)
}

// WriteToken writes token.
func (p Plugin) WriteToken() error {
	var err error

	p.Config.Remote, err = repo.WriteToken(
		p.Config.Remote,
		p.Netrc.Login,
		p.Netrc.Password,
	)

	return err
}

// HandleRemote adds the git remote if required.
func (p Plugin) HandleRemote(ctx context.Context) error {
	if p.Config.Remote != "" {
		if err := execute(repo.RemoteAdd(ctx, p.Config.RemoteName, p.Config.Remote)); err != nil {
			return err
		}
	}

	return nil
}

// HandlePath changes to a different directory if required
func (p Plugin) HandlePath() error {
	if p.Config.Path != "" {
		if err := os.Chdir(p.Config.Path); err != nil {
			return err
		}
	}

	return nil
}

// HandleCommit commits dirty changes if required.
func (p Plugin) HandleCommit(ctx context.Context) error {
	if p.Config.Commit {
		if err := execute(repo.Add(ctx)); err != nil {
			return err
		}

		if err := execute(repo.TestCleanTree(ctx)); err != nil {
			// changes to commit
			if err := execute(repo.ForceCommit(ctx, p.Config.CommitMessage, p.Config.NoVerify, p.Commit.Author.Name, p.Commit.Author.Email)); err != nil {
				return err
			}
		} else { // no changes
			if p.Config.EmptyCommit {
				// no changes but commit anyway
				if err := execute(repo.EmptyCommit(ctx, p.Config.CommitMessage, p.Config.NoVerify, p.Commit.Author.Name, p.Commit.Author.Email)); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

// HandleTag add tag if required.
func (p Plugin) HandleTag(ctx context.Context) error {
	if p.Config.Tag != "" {
		if err := execute(repo.Tag(ctx, p.Config.Tag)); err != nil {
			return err
		}
	}

	return nil
}

// HandlePush pushs the changes to the remote repo.
func (p Plugin) HandlePush(ctx context.Context) error {
	var (
		name       = p.Config.RemoteName
		local      = p.Config.LocalBranch
		branch     = p.Config.Branch
		force      = p.Config.Force
		followtags = p.Config.FollowTags
	)

	return execute(repo.RemotePushNamedBranch(ctx, name, local, branch, force, followtags))
}

// HanldeRebase pull rebases before pushing
func (p Plugin) HandleRebase(ctx context.Context) error {
	if p.Config.Rebase {
		var (
			name   = p.Config.RemoteName
			branch = p.Config.Branch
		)

		if err := execute(repo.RemotePullRebaseNamedBranch(ctx, name, branch)); err != nil {
			return err
		}
	}

	return nil
}

// HandleCleanup does eventually do some cleanup.
func (p Plugin) HandleCleanup(ctx context.Context) error {
	if p.Config.Remote != "" {
		if err := execute(repo.RemoteRemove(ctx, p.Config.RemoteName)); err != nil {
			return err
		}
	}

	return nil
}
