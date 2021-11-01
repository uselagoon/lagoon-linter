package lagoonyml

import (
	"fmt"
	"os"

	"sigs.k8s.io/yaml"
)

// Linter validates the given Lagoon struct.
type Linter func(*Lagoon) error

// Lint takes a file path, reads it, and applies `.lagoon.yml` lint policy to
// it. Lint returns an error of type ErrLint if it finds problems with the
// file, a regular error if something else went wrong, and nil if the
// `.lagoon.yml` is valid.
func Lint(path string, linters ...Linter) error {
	var l Lagoon
	l.Environments = make(map[string]Environment)

	rawYAML, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("couldn't read %v: %v", path, err)
	}

	// unmarshal the raw yaml into a map[string]interface{}
	var li map[string]interface{}
	err = yaml.Unmarshal(rawYAML, &li)
	if err != nil {
		return fmt.Errorf(".lagoon.yml configuration not valid for %v: %v", path, err)
	}

	// check each block for ability to be unmarshalled
	for key, block := range li {
		if b, ok := block.(map[string]interface{}); ok {
			switch key {
			case "environments":
				for env, config := range b {
					c, _ := yaml.Marshal(config)
					var le Environment
					err = yaml.Unmarshal(c, &le)
					if err != nil {
						fmt.Printf("Warning: .lagoon.yml configuration not valid for environment '%s': %v\n", env, err)
					}
					l.Environments[env] = le
				}
			case "production_routes":
				for env, config := range b {
					c, _ := yaml.Marshal(config)
					var le Environment
					err = yaml.Unmarshal(c, &le)
					if err != nil {
						fmt.Printf("Warning: .lagoon.yml configuration not valid for production_routes '%s': %v\n", env, err)
					}
					if env == "active" {
						l.ProductionRoutes.Active = le
					} else if env == "standby" {
						l.ProductionRoutes.Standby = le
					}
				}
			}
		}
	}

	// run the linter
	for _, linter := range linters {
		if err := linter(&l); err != nil {
			return &ErrLint{
				Detail: err.Error(),
			}
		}
	}
	return nil
}
