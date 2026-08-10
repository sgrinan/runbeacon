package tui

import "charm.land/bubbletea/v2"

// Update handles keyboard input and updates the model state.
func (model Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		model.viewport.SetHeight(msg.Height)
		model.viewport.SetWidth(msg.Width)

	case tea.KeyPressMsg:
		if model.selected != nil {
			var cmd tea.Cmd
			model.viewport, cmd = model.viewport.Update(msg)

			if msg.Code == tea.KeyEsc {
				model.selected = nil
			}

			return model, cmd
		}

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
				model.viewport.SetContent(model.runbookView())
			}
		}
	}

	return model, nil
}
