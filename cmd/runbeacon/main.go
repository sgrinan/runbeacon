package main

import (
	"fmt"
	"os"

	"github.com/sgrinan/runbeacon/internal/engine"
)

func main() {
	catalog, err := engine.LoadCatalog("runbooks")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	exitCode := run(os.Args[1:], os.Stdout, os.Stderr, catalog)

	os.Exit(exitCode)
}
