//go:build OMIT

package main

import (
	"io"
	"os"
	"strings"
)

type Upper struct {
	r io.Reader
}

func main() {
	s := strings.NewReader("phrase en minuscule")
	r := Upper{s}
	io.Copy(os.Stdout, &r)
}
