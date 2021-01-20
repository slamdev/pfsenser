package main

import (
	"github.com/slamdev/pfsenser/internal"
	"github.com/spf13/cobra/doc"
	"log"
)

func main() {
	err := doc.GenMarkdownTree(internal.RootCmd, "./docs")
	if err != nil {
		log.Fatal(err)
	}
}
