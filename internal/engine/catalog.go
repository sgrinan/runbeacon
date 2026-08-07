package engine

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/sgrinan/runbeacon/internal/models"
	"github.com/sgrinan/runbeacon/internal/parser"
	"github.com/sgrinan/runbeacon/internal/validator"
)

// Catalog contains the runbooks available to the application.
type Catalog struct {
	Runbooks []models.Runbook
}

// LoadCatalog loads and validates runbooks from a directory.
func LoadCatalog(directory string) (Catalog, error) {
	entries, err := os.ReadDir(directory)
	if err != nil {
		return Catalog{}, fmt.Errorf("read runbooks directory %q: %w", directory, err)
	}

	catalog := Catalog{}
	for _, entry := range entries {
		if entry.IsDir() || !isRunbookFile(entry.Name()) {
			continue
		}

		path := filepath.Join(directory, entry.Name())

		runbook, err := parser.Load(path)
		if err != nil {
			return Catalog{}, fmt.Errorf("load catalog runbook %q: %w", path, err)
		}

		if err := validator.Validate(runbook); err != nil {
			return Catalog{}, fmt.Errorf("validate runbook %q: %w", path, err)
		}

		catalog.Runbooks = append(catalog.Runbooks, runbook)
	}

	return catalog, nil
}

func isRunbookFile(name string) bool {
	return strings.HasSuffix(name, ".yaml") || strings.HasSuffix(name, ".yml")
}
