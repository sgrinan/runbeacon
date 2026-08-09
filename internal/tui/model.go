package tui

import (
	"charm.land/bubbletea/v2"
	"github.com/sgrinan/runbeacon/internal/engine"
	"github.com/sgrinan/runbeacon/internal/models"
)

type Model struct {
	catalog engine.Catalog
	query   string
	results []models.Runbook
	cursor  int
}

func NewModel(catalog engine.Catalog) Model {
	return Model{
		catalog: catalog,
		query:   "",
		results: catalog.Search(""),
		cursor:  0,
	}
}
func (model Model) Init() tea.Cmd {
	return nil
}

func (model Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
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

	}

	return model, nil
}

func (model Model) View() tea.View {
	return tea.NewView("")
}
