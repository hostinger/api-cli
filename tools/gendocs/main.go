package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/hostinger/api-cli/cmd"
	"github.com/spf13/cobra/doc"
)

// gendocs writes the markdown command reference (default, into ./docs) or,
// with -man <dir>, man pages for every command.
func main() {
	manDir := flag.String("man", "", "write man pages into this directory instead of markdown into ./docs")
	flag.Parse()

	var err error
	if *manDir != "" {
		if err = os.MkdirAll(*manDir, 0o755); err == nil {
			err = doc.GenManTree(cmd.RootCmd, &doc.GenManHeader{Section: "1", Manual: cmd.RootCmd.Short}, *manDir)
		}
	} else if err = os.MkdirAll("./docs", 0o755); err == nil {
		err = doc.GenMarkdownTree(cmd.RootCmd, "./docs")
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
