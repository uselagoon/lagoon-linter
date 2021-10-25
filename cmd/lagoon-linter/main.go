package main

import "github.com/alecthomas/kong"

// CLI represents the command-line interface.
type CLI struct {
	Validate ValidateCmd `kong:"cmd,default=1,help='(default) Validate the Lagoon YAML'"`
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
