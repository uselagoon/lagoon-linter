package main

import "github.com/amazeeio/lagoon-linter/internal/lagoonyml"

// ValidateCmd represents the validate command.
type ValidateCmd struct {
	LagoonYAML string `kong:"default='.lagoon.yml',help='Specify the Lagoon YAML file.'"`
}

// Run the validation of the Lagoon YAML.
func (cmd *ValidateCmd) Run() error {
	return lagoonyml.Lint(`.lagoon.yml`, lagoonyml.RouteAnnotation())
}
