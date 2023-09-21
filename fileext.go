package main

import "strings"

func hasTXTExtension(filename string) bool {
	return strings.HasSuffix(filename, ".txt")
}