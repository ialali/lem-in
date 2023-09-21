package errhand

import (
	"errors"
	"fmt"
	"strings"
)

func Hello() string {
	return "Hello, world!"
}

func HasTXTExtension(filename string) error {
	if strings.HasSuffix(filename, ".txt") {
		fmt.Println(filename, "has a .txt extension")
		return nil
	}
	return errors.New("not .txt file")
}
