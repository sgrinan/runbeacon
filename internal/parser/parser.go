package parser

import (
	"fmt"

	"github.com/sgrinan/runbeacon/internal/models"
	"go.yaml.in/yaml/v3"
)

func Parse(data []byte) (models.Runbook, error) {
	var runbook models.Runbook

	if err := yaml.Unmarshal(data, &runbook); err != nil {
		return models.Runbook{}, fmt.Errorf("parse runbook YAML: %w", err)
	}

	return runbook, nil
}
