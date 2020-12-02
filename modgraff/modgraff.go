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

type Config struct {
	DirPath      string
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

	modfstr, err := ioutil.ReadFile(modfname)
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

	return
}

func (e *Graff) readAll() (err error) {
	filepath.Walk(e.config.DirPath, func(path string, f os.FileInfo, err error) error {
		if !strings.HasSuffix(f.Name(), ".go") {
			return nil
		}

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

	filemodname := filepath.Join(node.Name.String())

	for _, i := range node.Imports {
		importname := strings.ReplaceAll(i.Path.Value, "\"", "")
		// fmt.Println(e.modpath, filemodname, importname)
		if !strings.Contains(importname, e.modpath) {
			break
		}
		if e.depmap[filemodname] == nil {
			e.depmap[filemodname] = map[string]bool{}
		}
		e.depmap[filemodname][importname] = true
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