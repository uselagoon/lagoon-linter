package main

import "github.com/alecthomas/kong"

var (
	date        string
	goVersion   string
	shortCommit string
	version     string
)

// CLI represents the command-line interface.
type CLI struct {
	Validate              ValidateCmd              `kong:"cmd,default=1,help='(default) Validate the Lagoon YAML'"`
	Version               VersionCmd               `kong:"cmd,help='Print version information'"`
	ValidateConfigMapJSON ValidateConfigMapJSONCmd `kong:"cmd,help='Validate the result of: kubectl get configmap -A -o json'"`
}

func main() {
	// parse CLI config
	cli := CLI{}
	kctx := kong.Parse(&cli,
		kong.UsageOnError(),
	)
	// execute CLI
	kctx.FatalIfErrorf(kctx.Run())
}
