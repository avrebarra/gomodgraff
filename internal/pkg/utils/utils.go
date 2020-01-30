package utils

import "strings"

func CleanStr(orig, removed string) string {
	return strings.ReplaceAll(orig, removed, "")
}
