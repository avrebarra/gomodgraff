package modgraff

import (
	"log"
	"os"
	"strings"
)

func verbose(enabled bool, v ...interface{}) {
	if !enabled {
		return
	}
	log.Println(v...)
}

func findModFile(dirpath string) (filename string) {
	dir, err := os.Open(dirpath)
	if err != nil {
		log.Fatalf("failed opening directory: %s", err)
	}
	defer dir.Close()

	files, _ := dir.Readdir(0)
	for _, f := range files {
		if f.IsDir() {
			continue
		}
		if strings.HasSuffix(f.Name(), ".mod") {
			return f.Name()
		}
	}

	return
}
