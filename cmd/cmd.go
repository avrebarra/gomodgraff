package cmd

import (
	"fmt"

	"github.com/leaanthony/clir"
)

var cmd *clir.Cli

func customBanner(cli *clir.Cli) string {
	return ``
}

func Initialize() {
	cmd = clir.NewCli("gomodgraff", "utility", "v2")
	// cmd.SetBannerFunction(customBanner)

	// default action
	path := "."
	onlyInternal := false
	verbose := false

	cmd.StringFlag("path", "golang project dir to analyze graph", &path)
	cmd.BoolFlag("only-internal", "only list internal dependencies", &onlyInternal)
	cmd.BoolFlag("verbose", "perform with process messages", &verbose)

	cmd.Action(func() error {
		subcmd := NewCommandMain(ConfigCommandMain{
			Path:         path,
			OnlyInternal: onlyInternal,
			Verbose:      verbose,
		})
		return subcmd.Run()
	})
}

func Execute() {
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
		return
	}
}
