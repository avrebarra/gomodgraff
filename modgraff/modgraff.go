package modgraff

import (
	"fmt"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"gopkg.in/go-playground/validator.v9"
)

const (
	internalModNotion = "&"
)

type Config struct {
	DirPath      string
	Verbose      bool
	OnlyInternal bool
}

type Graff struct {
	config  Config
	depmap  map[string]map[string]bool
	modpath string
}

func New(cfg Config) (g *Graff, err error) {
	if err := validator.New().Struct(cfg); err != nil {
		panic(err)
	}

	g = &Graff{
		config:  cfg,
		depmap:  map[string]map[string]bool{},
		modpath: "",
	}

	if err = g.setupModPath(); err != nil {
		return
	}

	if err = g.readAll(); err != nil {
		return
	}

	return
}

func (e *Graff) setupModPath() (err error) {
	modfname := findModFile(e.config.DirPath)
	if modfname == "" {
		err = fmt.Errorf("cannot find .mod file")
		return
	}

	modfstr, err := ioutil.ReadFile(filepath.Join(e.config.DirPath, modfname))
	if err != nil {
		err = fmt.Errorf("mod file read failure: %w", err)
		return
	}

	rgxmod := regexp.MustCompile(`module .+`)

	match := rgxmod.Find(modfstr)
	if len(match) == 0 {
		err = fmt.Errorf("mod file read failure")
		return
	}

	modpath := strings.SplitN(string(match), " ", 2)[1]
	if modpath == "" {
		err = fmt.Errorf("mod file read failure")
		return
	}

	e.modpath = modpath

	verbose(e.config.Verbose, "mod directory path", e.config.DirPath)
	verbose(e.config.Verbose, "mod name", e.modpath)

	return
}

func (e *Graff) readAll() (err error) {
	filepath.Walk(e.config.DirPath, func(path string, f os.FileInfo, err error) error {
		if !strings.HasSuffix(f.Name(), ".go") {
			return nil
		}
		if strings.HasSuffix(f.Name(), "_test.go") {
			return nil
		}

		verbose(e.config.Verbose, "found file", path)

		return e.Add(path)
	})

	return
}

func (e *Graff) Add(fpath string) (err error) {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, fpath, nil, parser.ParseComments)
	if err != nil {
		return
	}

	modname := filepath.Join(e.modpath, node.Name.String())
	finmodname := internalModNotion + strings.Replace(modname, e.modpath+"/", "", 1)
	verbose(e.config.Verbose, "  mod name", modname)
	verbose(e.config.Verbose, "  mod imports", len(node.Imports))

	for _, i := range node.Imports {
		importname := strings.ReplaceAll(i.Path.Value, "\"", "")
		verbose(e.config.Verbose, "      imported", importname)

		// seclude internal imports
		if e.config.OnlyInternal && !strings.Contains(importname, e.modpath) {
			continue
		}

		// shorten internal module names
		finimportmodname := importname
		if strings.Contains(importname, e.modpath+"/") {
			finimportmodname = internalModNotion + strings.Replace(importname, e.modpath+"/", "", 1)
		}

		// ensure store
		if e.depmap[finmodname] == nil {
			e.depmap[finmodname] = map[string]bool{}
		}
		e.depmap[finmodname][finimportmodname] = true
	}

	return
}

func (e *Graff) DotString() (dot string, err error) {
	dot += "digraph sample {\n"
	for pkgpath, defs := range e.depmap {
		for dep := range defs {
			dot += fmt.Sprintf(
				"\"%s\" -> \"%s\";\n",
				pkgpath, dep,
			)
		}
	}
	dot += "}"

	return
}
