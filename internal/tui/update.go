package tui

import tea "charm.land/bubbletea/v2"

// Update handles keyboard input and updates the model state.
func (model Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		model.list.SetSize(msg.Width, msg.Height)
		model.viewport.SetWidth(msg.Width)

		height := max(0, msg.Height-1)
		model.viewport.SetHeight(height)

	case tea.KeyPressMsg:
		if msg.String() == "ctrl+c" {
			return model, tea.Quit
		}

		if model.selected != nil && msg.Code == tea.KeyEsc {
			model.selected = nil
			return model, nil
		}

		if model.selected == nil && !model.list.SettingFilter() && msg.Code == tea.KeyEnter {
			item, ok := model.list.SelectedItem().(runbookItem)
			if ok {
				model.selected = &item.runbook
				model.viewport.SetContent(renderRunbook(item.runbook))
			}
			return model, nil
		}
	}

	if model.selected != nil {
		var cmd tea.Cmd
		model.viewport, cmd = model.viewport.Update(msg)
		return model, cmd
	}

	var cmd tea.Cmd
	model.list, cmd = model.list.Update(msg)
	return model, cmd
}
