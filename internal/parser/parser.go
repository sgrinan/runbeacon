package parser

import (
	"fmt"
	"os"

	"github.com/sgrinan/runbeacon/internal/models"
	"go.yaml.in/yaml/v3"
)

// Parse parses YAML data into a Runbook.
func Parse(data []byte) (models.Runbook, error) {
	var runbook models.Runbook

	if err := yaml.Unmarshal(data, &runbook); err != nil {
		return models.Runbook{}, fmt.Errorf("parse runbook YAML: %w", err)
	}

	return runbook, nil
}

// Load reads a runbook from a YAML file and parses it.
func Load(path string) (models.Runbook, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return models.Runbook{}, fmt.Errorf("read runbook file %q: %w", path, err)
	}

	return Parse(data)
}
