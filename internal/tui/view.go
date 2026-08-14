package tui

import (
	tea "charm.land/bubbletea/v2"

	"github.com/sgrinan/runbeacon/internal/render"
)

const runbookHelp = "↑/↓ scroll • esc back • ctrl+c quit"

// View renders the current model state.
func (model Model) View() tea.View {
	if model.selected != nil {
		content := model.viewport.View() + "\n" + runbookHelp
		return tea.NewView(content)
	}

	return tea.NewView(model.list.View())
}

func (model Model) runbookView() string {
	return render.Runbook(*model.selected)
}
