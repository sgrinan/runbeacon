package engine

import (
	"strings"

	"github.com/sgrinan/runbeacon/internal/models"
)

// Search returns runbooks matching the query.
func (catalog Catalog) Search(query string) []models.Runbook {
	query = normalize(query)

	var runbooks []models.Runbook
	for _, runbook := range catalog.Runbooks {
		if matchesRunbook(runbook, query) {
			runbooks = append(runbooks, runbook)
		}
	}

	return runbooks
}

func matchesRunbook(runbook models.Runbook, query string) bool {
	fields := []string{
		runbook.ID,
		runbook.Title,
		runbook.Description,
	}

	for _, symptom := range runbook.Symptoms {
		fields = append(fields, symptom.ID, symptom.Name)
		fields = append(fields, symptom.Keywords...)
	}

	for _, step := range runbook.Steps {
		fields = append(fields, step.Title, step.Description, step.Command)
	}

	for _, field := range fields {
		if strings.Contains(normalize(field), query) {
			return true
		}
	}
	return false
}

func normalize(value string) string {
	return strings.ToLower(strings.TrimSpace(value))
}
