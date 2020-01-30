package pack

import (
	"bufio"
	"os"
	"regexp"
	"strings"

	"github.com/shrotavre/gomodgraff/internal/pkg/utils"
)

func ReadModName(fpath string) (pkgname string) {
	f, _ := os.Open(fpath)
	defer f.Close()

	fscanner := bufio.NewScanner(f)
	for fscanner.Scan() {
		line := fscanner.Text()
		if line == "" || strings.HasPrefix(line, "//") {
			continue
		}
		return strings.Split(line, " ")[1]
	}

	return
}

func ReadImportedPkgs(fpath string) (pkgs []string) {
	f, _ := os.Open(fpath)
	defer f.Close()

	pkgs = []string{}
	rgx := regexp.MustCompile(`(".*")`)

	// states
	holdingLongImport := false

	fscanner := bufio.NewScanner(f)
	for fscanner.Scan() {
		line := fscanner.Text()

		if line == "import ()" {
			return
		}
		if line == "" || strings.HasPrefix(line, "//") {
			continue
		}
		if strings.HasPrefix(line, "//") {
			continue
		}

		if holdingLongImport {
			if line == ")" {
				holdingLongImport = false
				break // end search
			}

			name := strings.TrimSpace(line)
			if sp := strings.Split(name, " "); len(sp) == 2 {
				name = sp[1]
			}
			pkgs = append(pkgs, utils.CleanStr(name, "\""))
		} else {
			if !strings.HasPrefix(line, "import") {
				continue
			} else if strings.HasPrefix(line, "import (") {
				holdingLongImport = true
				continue
			} else {
				name := rgx.FindStringSubmatch(line)[0]
				pkgs = append(pkgs, utils.CleanStr(name, "\""))
			}
		}
	}

	return pkgs
}
