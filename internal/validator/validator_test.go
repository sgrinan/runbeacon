package validator

import (
	"strings"
	"testing"

	"github.com/sgrinan/runbeacon/internal/models"
)

func validRunbook() models.Runbook {
	return models.Runbook{
		ID:          "kubernetes.pod.crashloopbackoff",
		Title:       "Pod in CrashLoopBackOff",
		Description: "Diagnose a repeatedly crashing pod.",
		Symptoms: []models.Symptom{
			{
				ID:       "kubernetes.pod.crash_loop",
				Name:     "CrashLoopBackOff",
				Keywords: []string{"crash", "restart", "backoff"},
			},
		},
		Steps: []models.Step{
			{
				Title:       "Describe pod",
				Description: "Inspect pod events and container state.",
				Command:     "kubectl describe pod",
			},
		},
	}
}

func TestValidate(t *testing.T) {
	tests := []struct {
		name    string
		modify  func(*models.Runbook)
		wantErr string
	}{
		{
			name:    "valid runbook",
			modify:  nil,
			wantErr: "",
		},
		{
			name: "missing ID",
			modify: func(runbook *models.Runbook) {
				runbook.ID = "  "
			},
			wantErr: "runbook id is required",
		},
		{
			name: "missing title",
			modify: func(runbook *models.Runbook) {
				runbook.Title = "  "
			},
			wantErr: "runbook title is required",
		},
		{
			name: "missing description",
			modify: func(runbook *models.Runbook) {
				runbook.Description = "  "
			},
			wantErr: "runbook description is required",
		},
		{
			name: "no symptoms",
			modify: func(runbook *models.Runbook) {
				runbook.Symptoms = nil
			},
			wantErr: "runbook must contain at least one symptom",
		},
		{
			name: "no steps",
			modify: func(runbook *models.Runbook) {
				runbook.Steps = nil
			},
			wantErr: "runbook must contain at least one step",
		},
		{
			name: "symptom missing ID",
			modify: func(runbook *models.Runbook) {
				runbook.Symptoms[0].ID = "  "
			},
			wantErr: "runbook symptom 1: id is required",
		},
		{
			name: "symptom missing name",
			modify: func(runbook *models.Runbook) {
				runbook.Symptoms[0].Name = "  "
			},
			wantErr: "runbook symptom 1: name is required",
		},
		{
			name: "no keywords",
			modify: func(runbook *models.Runbook) {
				runbook.Symptoms[0].Keywords = nil
			},
			wantErr: "runbook symptom 1: must contain at least one keyword",
		},
		{
			name: "empty keyword",
			modify: func(runbook *models.Runbook) {
				runbook.Symptoms[0].Keywords = []string{"  "}
			},
			wantErr: "runbook symptom 1: keyword 1 is required",
		},
		{
			name: "step missing title",
			modify: func(runbook *models.Runbook) {
				runbook.Steps[0].Title = "  "
			},
			wantErr: "runbook step 1: title is required",
		},
		{
			name: "step missing description",
			modify: func(runbook *models.Runbook) {
				runbook.Steps[0].Description = "  "
			},
			wantErr: "runbook step 1: description is required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runbook := validRunbook()

			if tt.modify != nil {
				tt.modify(&runbook)
			}

			err := Validate(runbook)

			if tt.wantErr == "" {
				assertNoError(t, err, "Validate()")
				return
			}

			assertErrorContains(t, err, "Validate()", tt.wantErr)
		})
	}
}

func assertNoError(t *testing.T, err error, operation string) {
	t.Helper()

	if err != nil {
		t.Fatalf("%s: %v", operation, err)
	}
}

func assertErrorContains(t *testing.T, err error, operation string, want string) {
	t.Helper()

	if err == nil {
		t.Fatalf("%s: want an error", operation)
	}

	if !strings.Contains(err.Error(), want) {
		t.Fatalf("%s: error = %q, want it to contain %q", operation, err, want)
	}
}
