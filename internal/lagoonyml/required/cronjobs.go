package required

import (
	"fmt"
	"strings"
)

// ValidateCronjob returns an error if the command for the cronjob has any
// newlines, and nil otherwise.
func ValidateCronjob(c *LagoonCronjob) error {
	if strings.Contains(strings.TrimSpace(c.Command), "\n") || strings.HasSuffix(c.Command, "\n") {
		return fmt.Errorf("invalid cronjob, multiline commands are not supported: %q",
			c.Command)
	}

	return nil
}

// Cronjobs checks for valid environment cronjobs.
func Cronjobs(l *Lagoon) error {
	for eName, e := range l.Environments {
		for _, lagoonCronjob := range e.Cronjobs {
			if err := ValidateCronjob(&lagoonCronjob); err != nil {
				return fmt.Errorf("environment %s: %v", eName, err)
			}
		}
	}

	return nil
}
