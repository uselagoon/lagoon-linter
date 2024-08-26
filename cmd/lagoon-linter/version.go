package main

import (
	"fmt"
	"runtime"
)

// VersionCmd represents the version command.
type VersionCmd struct{}

// Run the version command to print version information.
func (cmd *VersionCmd) Run() error {
	fmt.Printf("lagoon-linter %v (%v) compiled with %v on %v\n", version,
		shortCommit, runtime.Version(), date)
	return nil
}
