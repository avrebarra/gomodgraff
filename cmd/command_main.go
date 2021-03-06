package cmd

import (
	"os"
	"path/filepath"

	"github.com/shrotavre/gomodgraff/modgraff"
	"gopkg.in/go-playground/validator.v9"
)

type ConfigCommandMain struct {
	Path         string `validate:"required"`
	OnlyInternal bool
	Verbose      bool
}

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
	normpath, err := filepath.Abs(c.config.Path)
	if err != nil {
		return
	}

	g, err := modgraff.New(modgraff.Config{
		DirPath:      normpath,
		OnlyInternal: c.config.OnlyInternal,
		Verbose:      c.config.Verbose,
	})
	if err != nil {
		return
	}

	dotstr, err := g.DotString()
	if err != nil {
		return
	}

	os.Stdout.Write([]byte(dotstr))

	return nil
}
