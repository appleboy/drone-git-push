package main

import (
	"context"
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
