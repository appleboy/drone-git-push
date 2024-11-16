package repo

import (
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
