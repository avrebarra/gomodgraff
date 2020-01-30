package main

import (
	"fmt"
	"strings"

	"github.com/shrotavre/gomodgraff/internal/pkg/pack"
)

func PrintDOT(m DependencyMap) {
	fmt.Println("digraph sample {")
	for pkgpath, deps := range m {
		for dep := range deps {

			fmt.Printf(
				"\"%s\" -> \"%s\";\n",
				pkgpath, dep,
			)
		}
	}
	fmt.Println("}")
}

func GetRelPkgName(fpath string) string {
	relpath := pack.GetRelPath(fpath)
	modname := ""
	if pack.IsDir(fpath) {
		files := pack.FindGoFiles(fpath)
		if len(files) == 0 {
			return ""
		}

		modname = pack.ReadModName(files[0])
	} else {
		modname = pack.ReadModName(fpath)
	}

	if relpath == "" {
		return modname
	}
	if strings.HasSuffix(relpath, modname) {
		return relpath
	}

	return relpath + "/" + modname
}
