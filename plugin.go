package main

import (
	"github.com/drone-plugins/drone-git-push/repo"
)

type (
	Netrc struct {
		Machine  string
		Login    string
		Password string
	}

	Commit struct {
		Author Author
	}

	Author struct {
		Name  string
		Email string
	}

	Config struct {
		Key         string
		Remote      string
		RemoteName  string
		Branch      string
		LocalBranch string
		Force       bool
		SkipVerify  bool
		Commit      bool
	}

	Plugin struct {
		Netrc  Netrc
		Commit Commit
		Config Config
	}
)

// Exec starts the plugin execution.
func (p Plugin) Exec() error {
	if err := p.WriteConfig(); err != nil {
		return err
	}

	if err := p.WriteKey(); err != nil {
		return err
	}

	if err := p.WriteNetrc(); err != nil {
		return err
	}

	if err := p.HandleCommit(); err != nil {
		return err
	}

	if err := p.HandleRemote(); err != nil {
		return err
	}

	if err := p.HandlePush(); err != nil {
		return err
	}

	if err := p.HandleCleanup(); err != nil {
		return err
	}

	return nil
}

// WriteConfig writes all required configurations.
func (p Plugin) WriteConfig() error {
	if err := repo.GlobalName(p.Commit.Author.Name).Run(); err != nil {
		return err
	}

	if err := repo.GlobalUser(p.Commit.Author.Email).Run(); err != nil {
		return err
	}

	if p.Config.SkipVerify {
		if err := repo.SkipVerify().Run(); err != nil {
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

// HandleRemote adds the git remote if required.
func (p Plugin) HandleRemote() error {
	if p.Config.Remote != "" {
		if err := execute(repo.RemoteAdd(p.Config.RemoteName, p.Config.Remote)); err != nil {
			return err
		}
	}

	return nil
}

// HandleCommit commits dirty changes if required.
func (p Plugin) HandleCommit() error {
	if p.Config.Commit {
		if err := execute(repo.ForceAdd()); err != nil {
			return err
		}

		if err := execute(repo.ForceCommit()); err != nil {
			return err
		}
	}

	return nil
}

// HandlePush pushs the changes to the remote repo.
func (p Plugin) HandlePush() error {
	var (
		name   = p.Config.RemoteName
		local  = p.Config.LocalBranch
		branch = p.Config.Branch
		force  = p.Config.Force
	)

	if err := execute(repo.RemotePushNamedBranch(name, local, branch, force)); err != nil {
		return err
	}

	return nil
}

// HandleCleanup does eventually do some cleanup.
func (p Plugin) HandleCleanup() error {
	if p.Config.Remote != "" {
		if err := execute(repo.RemoteRemove(p.Config.RemoteName)); err != nil {
			return err
		}
	}

	return nil
}
