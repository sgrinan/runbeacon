package tui

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"

	"github.com/sgrinan/runbeacon/internal/render"
)

// View renders the current model state.
func (model Model) View() tea.View {
	if model.selected != nil {
		return tea.NewView(model.runbookView())
	}

	return tea.NewView(model.searchView())
}

func (model Model) searchView() string {
	var builder strings.Builder

	fmt.Fprintf(&builder, "Search: %s\n\n", model.query)

	for i, runbook := range model.results {
		if i == model.cursor {
			fmt.Fprintf(&builder, ">  %s\n", runbook.Title)
		} else {
			fmt.Fprintf(&builder, "   %s\n", runbook.Title)
		}
	}

	return builder.String()
}

func (model Model) runbookView() string {
	return render.Runbook(*model.selected)
}
