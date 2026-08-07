package parser

import (
	"reflect"
	"strings"
	"testing"

	"github.com/sgrinan/runbeacon/internal/models"
)

const validRunbookYAML = `
id: kubernetes.pod.crashloopbackoff
title: Pod in CrashLoopBackOff
description: Diagnose a repeatedly crashing pod.

symptoms:
  - id: kubernetes.pod.crash_loop
    name: CrashLoopBackOff
    keywords:
      - crash
      - restart
      - backoff

steps:
  - title: Describe pod
    description: Inspect pod events and container state.
    command: kubectl describe pod
`

const invalidRunbookYAML = `
id: kubernetes.pod.crashloopbackoff
title: [
`

// TestParseValidYAML
func TestParseValidYAML(t *testing.T) {

	want := models.Runbook{
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

	got, err := Parse([]byte(validRunbookYAML))
	assertNoError(t, err, "Parse()")

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Parse() = %#v, want %#v", got, want)
	}
}

// TestParseInvalidYAML
func TestParseInvalidYAML(t *testing.T) {
	_, err := Parse([]byte(invalidRunbookYAML))
	assertErrorContains(t, err, "Parse()", "parse runbook YAML")
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
