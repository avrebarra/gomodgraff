package cmd

import (
	"fmt"

	"github.com/shrotavre/gomodgraff/modgraff"
	"gopkg.in/go-playground/validator.v9"
)

type ConfigCommandMain struct{}

type CommandMain struct {
	config ConfigCommandMain
}

func NewCommandMain(cfg ConfigCommandMain) CommandMain {
	if err := validator.New().Struct(cfg); err != nil {
		panic(err)
	}

	cmd := CommandMain{config: cfg}

	return cmd
}

func (c *CommandMain) Run() (err error) {
	g, err := modgraff.New(modgraff.Config{
		DirPath:      ".",
		OnlyInternal: true,
	})
	if err != nil {
		return
	}

	fmt.Println(g.DotString())

	return nil
}
