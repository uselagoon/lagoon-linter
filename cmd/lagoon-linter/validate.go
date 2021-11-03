package main

import (
	"fmt"
	"os"

	"github.com/uselagoon/lagoon-linter/internal/lagoonyml"
)

// ValidateCmd represents the validate command.
type ValidateCmd struct {
	LagoonYAML string `kong:"default='.lagoon.yml',help='Specify the Lagoon YAML file.'"`
}

// Run the validation of the Lagoon YAML.
func (cmd *ValidateCmd) Run() error {
	rawYAML, err := os.ReadFile(cmd.LagoonYAML)
	if err != nil {
		return fmt.Errorf("couldn't read %v: %v", cmd.LagoonYAML, err)
	}
	return lagoonyml.Lint(rawYAML, lagoonyml.RouteAnnotation())
}
