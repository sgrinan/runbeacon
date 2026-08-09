package tui

import (
	tea "charm.land/bubbletea/v2"

	"github.com/sgrinan/runbeacon/internal/engine"
	"github.com/sgrinan/runbeacon/internal/models"
)

// Model holds the state of the interactive runbook search.
type Model struct {
	catalog engine.Catalog
	query   string
	results []models.Runbook
	cursor  int
	selected *models.Runbook
}

// NewModel creates a model with all runbooks initially visible.
func NewModel(catalog engine.Catalog) Model {
	return Model{
		catalog: catalog,
		query:   "",
		results: catalog.Search(""),
		cursor:  0,
		selected: nil,
	}
}

// Init initializes the Bubble Tea model.
func (model Model) Init() tea.Cmd {
	return nil
}

// Update handles keyboard input and updates the model state.
func (model Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		if msg.String() == "ctrl+c" {
			return model, tea.Quit
		}

		if msg.Code == tea.KeyDown {
			if model.cursor < len(model.results)-1 {
				model.cursor++
			}
		}

		if msg.Code == tea.KeyUp {
			if model.cursor > 0 {
				model.cursor--
			}
		}

		if msg.Text != "" {
			model.query += msg.Text
			model.results = model.catalog.Search(model.query)
			model.cursor = 0
		}

		if msg.Code == tea.KeyBackspace {
			if len(model.query) > 0 {
				model.query = model.query[:len(model.query)-1]
				model.results = model.catalog.Search(model.query)
				model.cursor = 0
			}
		}

		if msg.Code == tea.KeyEnter {
			if len(model.results) > 0 {
				model.selected = &model.results[model.cursor]
			}
		}
	}

	return model, nil
}

// View renders the current model state.
func (model Model) View() tea.View {
	return tea.NewView("")
}
