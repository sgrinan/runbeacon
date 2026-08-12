package tui

import (
	tea "charm.land/bubbletea/v2"

	"github.com/sgrinan/runbeacon/internal/render"
)

// View renders the current model state.
func (model Model) View() tea.View {
	if model.selected != nil {
		return tea.NewView(model.viewport.View())
	}

	return tea.NewView(model.list.View())
}

func (model Model) runbookView() string {
	return render.Runbook(*model.selected)
}
