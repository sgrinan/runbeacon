package tui

import "github.com/sgrinan/runbeacon/internal/models"

type runbookItem struct {
	runbook models.Runbook
}

func (item runbookItem) Title() string {
	return item.runbook.Title
}

func (item runbookItem) Description() string {
	return item.runbook.Description
}

func (item runbookItem) FilterValue() string {
	return item.runbook.Title
}
