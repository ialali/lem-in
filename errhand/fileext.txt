package lemin

import (
	"strings"
)

func hasTXTExtension(filename string) bool {
	return strings.HasSuffix(filename, ".txt")
}