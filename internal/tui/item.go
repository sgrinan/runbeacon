package tui

import (
	"strings"

	"github.com/sgrinan/runbeacon/internal/models"
)

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
	var terms []string

	terms = append(terms, item.runbook.Title)
	terms = append(terms, item.runbook.ID)
	terms = append(terms, item.runbook.Description)

	for _, symptom := range item.runbook.Symptoms {
		terms = append(terms, symptom.Name)
		terms = append(terms, symptom.Keywords...)
	}

	return strings.Join(terms, " ")
}
