package main

import (
	"github.com/jessevdk/go-flags"
	"github.com/kangaechu/waituntil/pkg/waituntil"
	"os"
)

var opts waituntil.CmdOpts

func main() {
	parser := flags.NewParser(&opts, flags.Default)
	_, err := flags.Parse(&opts)
	if err != nil {
		parser.WriteHelp(os.Stdout)
		os.Exit(1)
	}
	waituntil.Run(opts)
}
