package lagoonyml_test

import (
	"testing"

	"github.com/uselagoon/lagoon-linter/internal/lagoonyml"
)

func TestLint(t *testing.T) {
	var testCases = map[string]struct {
		input string
		valid bool
	}{
		"valid.0.lagoon.yml": {
			input: "testdata/valid.0.lagoon.yml",
			valid: true,
		},
		"valid.1.lagoon.yml": {
			input: "testdata/valid.1.lagoon.yml",
			valid: true,
		},
		"invalid.0.lagoon.yml": {
			input: "testdata/invalid.0.lagoon.yml",
			valid: false,
		},
		"invalid.1.lagoon.yml": {
			input: "testdata/invalid.1.lagoon.yml",
			valid: false,
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(tt *testing.T) {
			err := lagoonyml.Lint(tc.input, lagoonyml.RouteAnnotation())
			if tc.valid {
				if err != nil {
					tt.Fatalf("unexpected error %v", err)
				}
			} else {
				tt.Log(err)
				if err == nil {
					tt.Fatalf("expected error, but got nil")
				}
			}
		})
	}
}
