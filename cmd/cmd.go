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
	cmd.Action(func() error {
		subcmd := NewCommandMain(ConfigCommandMain{})
		return subcmd.Run()
	})
}

func Execute() {
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
		return
	}
}
