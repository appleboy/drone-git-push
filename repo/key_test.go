package repo

import "testing"

func TestWriteToken(t *testing.T) {
	type args struct {
		remote   string
		login    string
		password string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "missing token",
			args: args{
				remote: "https://github.com/foo/bar.git",
				login:  "foo",
			},
			want:    "https://github.com/foo/bar.git",
			wantErr: false,
		},
		{
			name: "add token",
			args: args{
				remote:   "https://github.com/foo/bar.git",
				login:    "foo",
				password: "bar",
			},
			want:    "https://foo:bar@github.com/foo/bar.git",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := WriteToken(tt.args.remote, tt.args.login, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("WriteToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("WriteToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
