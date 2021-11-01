package lagoonyml

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"sigs.k8s.io/yaml"
)

// Linter validates the given Lagoon struct.
type Linter func(*Lagoon) error

func lint(rawYAML []byte, linters []Linter) error {
	var l Lagoon
	yamlErr := yaml.Unmarshal(rawYAML, &l)
	if yamlErr != nil {
		// try to extract more detail from the Unmarshal error
		rawJSON, err := yaml.YAMLToJSON(rawYAML)
		if err != nil {
			// can't even convert this so return the original error
			return yamlErr
		}
		// json.Unmarshal returns richer errors than yaml.Unmarshal, which helps to
		// diagnose exactly what went wrong
		err = json.Unmarshal(rawJSON, &l)
		var jTypeErr *json.UnmarshalTypeError
		if errors.As(err, &jTypeErr) {
			// this rawJSON slice will be partial JSON, but JSONToYAML ignores junk
			// at the end of the snippet.
			badYAML, err := yaml.JSONToYAML(rawJSON[jTypeErr.Offset:])
			if err != nil {
				// can't convert this snippet, so return the original error
				return yamlErr
			}
			return fmt.Errorf(
				"couldn't unmarshal YAML: %v.\nThere appears to be invalid YAML in the `%s` field:\n\n%s",
				yamlErr, jTypeErr.Field, badYAML)
		}
		// this isn't a json.UnmarshalTypeError, so just return the original error
		return yamlErr
	}
	for _, linter := range linters {
		if err := linter(&l); err != nil {
			return &ErrLint{
				Detail: err.Error(),
			}
		}
	}
	return nil
}

// LintFile takes a file path, reads it, and applies `.lagoon.yml` lint policy
// to it. LintFile returns an error of type ErrLint if it finds problems with
// the file, a regular error if something else went wrong, and nil if the
// `.lagoon.yml` is valid.
func LintFile(path string, linters ...Linter) error {
	rawYAML, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("couldn't read %v: %v", path, err)
	}
	if err = lint(rawYAML, linters); err != nil {
		return fmt.Errorf("couldn't validate: %v: %v", path, err)
	}
	return nil
}

// LintYAML takes a byte slice containing raw YAML and applies `.lagoon.yml`
// lint policy to it. LintYAML returns an error of type ErrLint if it finds
// problems with the YAML, a regular error if something else went wrong, and
// nil if the YAML is valid.
func LintYAML(rawYAML []byte, linters ...Linter) error {
	return lint(rawYAML, linters)
}
