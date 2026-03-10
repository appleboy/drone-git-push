package repo

import (
	"context"
	"testing"
)

func TestSanitizeInput(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "valid input with alphanumeric characters",
			input: "feature-branch",
			want:  "feature-branch",
		},
		{
			name:  "valid input with dots and slashes",
			input: "release/1.0.0",
			want:  "release/1.0.0",
		},
		{
			name:  "invalid input with spaces",
			input: "invalid branch",
			want:  "",
		},
		{
			name:  "invalid input with special characters",
			input: "invalid@branch!",
			want:  "",
		},
		{
			name:  "empty input",
			input: "",
			want:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sanitizeInput(tt.input); got != tt.want {
				t.Errorf("sanitizeInput() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoteSetURL(t *testing.T) {
	cmd := RemoteSetURL(context.Background(), "origin", "git@github.com:user/repo.git")
	args := cmd.Args
	expected := []string{"git", "remote", "set-url", "origin", "git@github.com:user/repo.git"}
	if len(args) != len(expected) {
		t.Fatalf("expected %d args, got %d", len(expected), len(args))
	}
	for i, arg := range args {
		if arg != expected[i] {
			t.Errorf("arg[%d] = %q, want %q", i, arg, expected[i])
		}
	}
}
