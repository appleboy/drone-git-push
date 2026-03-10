package main

import (
	"context"
	"os"
	"os/exec"
	"testing"
)

func TestPlugin_HandleRemote(t *testing.T) {
	type fields struct {
		Netrc  Netrc
		Commit Commit
		Config Config
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "vaild git URL",
			fields: fields{
				Config: Config{
					RemoteName: "deploy",
					Remote:     "git@github.com:appleboy/drone-git-push.git",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Plugin{
				Netrc:  tt.fields.Netrc,
				Commit: tt.fields.Commit,
				Config: tt.fields.Config,
			}
			if err := p.HandleRemote(context.Background()); (err != nil) != tt.wantErr {
				t.Errorf("Plugin.HandleRemote() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPlugin_HandleRemote_ExistingRemote(t *testing.T) {
	// Create a temporary git repo
	tmpDir, err := os.MkdirTemp("", "drone-git-push-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	// Save current dir and change to temp dir
	origDir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		_ = os.Chdir(origDir)
	}()

	if err := os.Chdir(tmpDir); err != nil {
		t.Fatal(err)
	}

	// Initialize a git repo and add a remote
	ctx := context.Background()
	if out, err := exec.CommandContext(
		ctx, "git", "init",
	).CombinedOutput(); err != nil {
		t.Fatalf("git init failed: %s, %v", out, err)
	}
	if out, err := exec.CommandContext(
		ctx, "git", "remote", "add", "origin",
		"git@github.com:old/repo.git",
	).CombinedOutput(); err != nil {
		t.Fatalf("git remote add failed: %s, %v", out, err)
	}

	// HandleRemote should succeed even though "origin" already exists
	p := Plugin{
		Config: Config{
			RemoteName: "origin",
			Remote:     "git@github.com:new/repo.git",
		},
	}
	if err := p.HandleRemote(ctx); err != nil {
		t.Errorf("HandleRemote() with existing remote should not fail, got: %v", err)
	}

	// Verify the remote URL was updated
	out, err := exec.CommandContext(
		ctx, "git", "remote", "get-url", "origin",
	).Output()
	if err != nil {
		t.Fatal(err)
	}
	got := string(out)
	want := "git@github.com:new/repo.git\n"
	if got != want {
		t.Errorf("remote URL = %q, want %q", got, want)
	}
}
