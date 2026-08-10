package tui

import (
	"charm.land/bubbles/v2/viewport"
	"charm.land/bubbletea/v2"

	"github.com/sgrinan/runbeacon/internal/engine"
	"github.com/sgrinan/runbeacon/internal/models"
)

// Model holds the state of the interactive runbook search.
type Model struct {
	catalog  engine.Catalog
	query    string
	results  []models.Runbook
	cursor   int
	selected *models.Runbook
	viewport viewport.Model
}

// NewModel creates a model with all runbooks initially visible.
func NewModel(catalog engine.Catalog) Model {
	return Model{
		catalog:  catalog,
		query:    "",
		results:  catalog.Search(""),
		cursor:   0,
		selected: nil,
		viewport: viewport.New(),
	}
}

// Init initializes the Bubble Tea model.
func (model Model) Init() tea.Cmd {
	return nil
}
