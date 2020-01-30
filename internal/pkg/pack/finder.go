package pack

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

func FindModFile(path string) (modfile string) {
	dir, err := os.Open(path)
	if err != nil {
		log.Fatalf("failed opening directory: %s", err)
	}
	defer dir.Close()

	files, _ := dir.Readdir(0)
	for _, f := range files {
		if f.IsDir() {
			continue
		}

		if HasExtension(f, "mod") {
			return f.Name()
		}
	}

	return
}

func FindGoFiles(path string) (gofiles []string) {
	gofiles = []string{}

	if !IsDir(path) {
		return
	}

	dir, err := os.Open(path)
	if err != nil {
		log.Fatalf("failed opening directory: %s", err)
	}
	defer dir.Close()

	files, _ := dir.Readdir(0)
	for _, f := range files {
		if f.IsDir() {
			continue
		}

		if HasExtension(f, "go") && !strings.HasSuffix(f.Name(), "_test.go") {
			gofiles = append(gofiles, f.Name())
		}
	}

	return
}

func FindGoFilesDeep(path string) (gofiles []string) {
	gofiles = []string{}

	if !IsDir(path) {
		return
	}

	filepath.Walk(path, func(path string, fi os.FileInfo, err error) error {
		if HasExtension(fi, "go") && !strings.HasSuffix(fi.Name(), "_test.go") {
			gofiles = append(gofiles, path)
		}

		return nil
	})

	return
}

func GetRelPath(fpath string) (p string) {
	if IsDir(fpath) {
		p = fpath
	} else {
		pathParts := strings.Split(fpath, "/")
		p = strings.Join(pathParts[0:len(pathParts)-1], "/")
	}

	return
}
