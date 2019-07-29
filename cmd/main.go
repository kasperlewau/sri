package main

import (
	"os"

	"github.com/kasperlewau/sri"
)

func main() {
	sri.Hash(os.Stdin, os.Stdout)
}
