package pack

import (
	"fmt"
	"os"
	"strings"
)

func IsDir(fpath string) bool {
	fi, err := os.Stat(fpath)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return fi.IsDir()
}

func HasExtension(fi os.FileInfo, ext string) bool {
	return strings.HasSuffix(fi.Name(), "."+ext)
}
