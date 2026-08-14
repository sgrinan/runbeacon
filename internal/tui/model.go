package tui

import (
	"charm.land/bubbles/v2/list"
	"charm.land/bubbles/v2/viewport"
	tea "charm.land/bubbletea/v2"

	"github.com/sgrinan/runbeacon/internal/engine"
	"github.com/sgrinan/runbeacon/internal/models"
)

// Model holds the state of the interactive runbook search.
type Model struct {
	selected *models.Runbook
	list     list.Model
	viewport viewport.Model
}

// NewModel creates a model with all runbooks initially visible.
func NewModel(catalog engine.Catalog) Model {
	items := make([]list.Item, 0, len(catalog.Runbooks))

	for _, runbook := range catalog.Runbooks {
		items = append(items, runbookItem{
			runbook: runbook,
		})
	}

	runbookList := list.New(items, list.NewDefaultDelegate(), 0, 0)
	runbookList.Filter = filterRunbooks

	return Model{
		selected: nil,
		list:     runbookList,
		viewport: viewport.New(),
	}
}

// Init initializes the Bubble Tea model.
func (model Model) Init() tea.Cmd {
	return nil
}
