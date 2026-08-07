package validator

import (
	"errors"
	"fmt"
	"strings"

	"github.com/sgrinan/runbeacon/internal/models"
)

// Validate checks that a runbook contains all required fields.
func Validate(runbook models.Runbook) error {
	if strings.TrimSpace(runbook.ID) == "" {
		return errors.New("runbook id is required")
	}

	if strings.TrimSpace(runbook.Title) == "" {
		return errors.New("runbook title is required")
	}

	if strings.TrimSpace(runbook.Description) == "" {
		return errors.New("runbook description is required")
	}

	if len(runbook.Symptoms) == 0 {
		return errors.New("runbook must contain at least one symptom")
	}

	for i, symptom := range runbook.Symptoms {
		if err := validateSymptom(symptom); err != nil {
			return fmt.Errorf("runbook symptom %d: %w", i+1, err)
		}
	}

	if len(runbook.Steps) == 0 {
		return errors.New("runbook must contain at least one step")
	}

	for i, step := range runbook.Steps {
		if err := validateStep(step); err != nil {
			return fmt.Errorf("runbook step %d: %w", i+1, err)
		}
	}

	return nil
}

func validateSymptom(symptom models.Symptom) error {
	if strings.TrimSpace(symptom.ID) == "" {
		return errors.New("id is required")
	}

	if strings.TrimSpace(symptom.Name) == "" {
		return errors.New("name is required")
	}

	if len(symptom.Keywords) == 0 {
		return errors.New("must contain at least one keyword")
	}

	for i, keyword := range symptom.Keywords {
		if strings.TrimSpace(keyword) == "" {
			return fmt.Errorf("keyword %d is required", i+1)
		}
	}

	return nil
}

func validateStep(step models.Step) error {
	if strings.TrimSpace(step.Title) == "" {
		return errors.New("title is required")
	}

	if strings.TrimSpace(step.Description) == "" {
		return errors.New("description is required")
	}

	return nil
}
