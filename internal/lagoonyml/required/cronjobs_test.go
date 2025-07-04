package required_test

import (
	"testing"

	"github.com/uselagoon/lagoon-linter/internal/lagoonyml/required"
	"sigs.k8s.io/yaml"
)

// All the possible YAML incantations for introducing a newline into a string.
// https://yaml-multiline.info/
var multiline = `environments:
  main:
    cronjobs:
      - name: flow scalar plain
        command: multiline

          command
      - name: flow scalar single quoted
        command: 'multiline

          command'
      - name: flow scalar double quoted newline
        command: "multiline

          command"
      - name: flow scalar double quoted escaped
        command: "multiline\ncommand"
      - name: multiline block literal clipped
        command: |
          multiline
          command

      - name: block scalar literal stripped
        command: |-
          multiline
          command

      - name: block scalar literal keep
        command: |+
          multiline
          command

      - name: block scalar folded clipped
        command: >
          multiline

          command

      - name: block scalar folded stripped
        command: >-
          multiline

          command

      - name: block scalar folded keep
        command: >+
          multiline

          command

`

// Strings that may appear to have newlines but don't.
var singleline = `environments:
  main:
    cronjobs:
      - name: flow scalar plain 1
        command: singleline
          command
      - name: flow scalar plain 2
        command: singleline command
      - name: flow scalar plain 3
        command: singleline\ncommand
      - name: flow scalar single quoted 1
        command: 'singleline
          command'
      - name: flow scalar single quoted 2
        command: 'singleline command'
      - name: flow scalar single quoted 3
        command: 'singleline\ncommand'
      - name: flow scalar double quoted 1
        command: "singleline
          command"
      - name: flow scalar double quoted 2
        command: "singleline command"
      - name: flow scalar double quoted 3
        command: "singleline\
          command"
      - name: block scalar literal stripped
        command: |-
          singleline command

      - name: block scalar folded stripped 1
        command: >-
          singleline
          command
      - name: block scalar folded stripped 2
        command: >-
          singleline command

`

func TestMultilineCommand(t *testing.T) {
	var l required.Lagoon
	if err := yaml.Unmarshal([]byte(multiline), &l); err != nil {
		t.Fatalf("couldn't unmarshal YAML: %v", err)
	}

	for _, e := range l.Environments {
		for _, lagoonCronjob := range e.Cronjobs {
			t.Run(lagoonCronjob.Name, func(tt *testing.T) {
				err := required.ValidateCronjob(&lagoonCronjob)

				tt.Log(err)
				if err == nil {
					tt.Fatalf("expected error, but got nil")
				}
			})
		}
	}

}

func TestSinglelineCommand(t *testing.T) {
	var l required.Lagoon
	if err := yaml.Unmarshal([]byte(singleline), &l); err != nil {
		t.Fatalf("couldn't unmarshal YAML: %v", err)
	}

	for _, e := range l.Environments {
		for _, lagoonCronjob := range e.Cronjobs {
			t.Run(lagoonCronjob.Name, func(tt *testing.T) {
				err := required.ValidateCronjob(&lagoonCronjob)

				if err != nil {
					tt.Fatalf("unexpected error %v", err)
				}
			})
		}
	}

}
