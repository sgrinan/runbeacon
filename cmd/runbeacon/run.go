package main

import (
	"fmt"
	"io"

	tea "charm.land/bubbletea/v2"

	"github.com/sgrinan/runbeacon/internal/engine"
	"github.com/sgrinan/runbeacon/internal/tui"
)

func run(args []string, stdout io.Writer, stderr io.Writer, catalog engine.Catalog) int {
	if len(args) == 0 {
		return runTUI(stdout, stderr, catalog)
	}

	switch args[0] {
	case "search":
		return runSearch(args[1:], stdout, stderr, catalog)
	default:
		fmt.Fprintf(stderr, "unknown command: %s\n", args[0])
		return 2
	}
}

func runTUI(stdout io.Writer, stderr io.Writer, catalog engine.Catalog) int {
	model := tui.NewModel(catalog)

	program := tea.NewProgram(
		model,
		tea.WithOutput(stdout),
	)

	if _, err := program.Run(); err != nil {
		fmt.Fprintf(stderr, "tui: %v\n", err)
		return 1
	}

	return 0
}

func runSearch(args []string, stdout io.Writer, stderr io.Writer, catalog engine.Catalog) int {
	if len(args) == 0 {
		fmt.Fprintln(stderr, "usage: runbeacon search <query>")
		return 2
	}

	runbooks := catalog.Search(args[0])

	for _, runbook := range runbooks {
		fmt.Fprintf(stdout, "%s\t%s\n", runbook.ID, runbook.Title)
	}

	return 0
}
