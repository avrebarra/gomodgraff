package main

import (
	"flag"
	"path/filepath"
	"strings"

	"github.com/shrotavre/gomodgraff/internal/pkg/pack"
	"github.com/shrotavre/gomodgraff/internal/pkg/utils"
)

type DependencyMap map[string]map[string]bool

func main() {
	dirpath := flag.String("dir", ".", "target directory to draw dependency graph from")
	internalOnly := flag.Bool("internal", false, "show only internal packages")

	flag.Parse()

	depsmapping := DependencyMap{}
	modscope := pack.ReadModName(filepath.Join(*dirpath, pack.FindModFile(*dirpath)))
	gofpaths := pack.FindGoFilesDeep(*dirpath)

	// build dependency mapping
	for _, gofpath := range gofpaths {
		pkgpath := utils.CleanStr(GetRelPkgName(gofpath), *dirpath+"/")
		if pkgpath == "" {
			continue
		}

		pkgimports := pack.ReadImportedPkgs(gofpath)
		if depsmapping[pkgpath] == nil {
			depsmapping[pkgpath] = make(map[string]bool)
		}

		for _, pkgi := range pkgimports {
			if *internalOnly && !strings.HasPrefix(pkgi, string(modscope)) {
				continue
			}
			pkgi = utils.CleanStr(pkgi, modscope+"/")
			depsmapping[pkgpath][pkgi] = true
		}
	}

	PrintDOT(depsmapping)
}
