package parser

import (
	"os"
	"reflect"
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

// TestParseValidYAML
func TestParseValidYAML(t *testing.T) {
	data, err := os.ReadFile("testdata/valid_runbook.yaml")
	assertNoError(t, err, "ReadFile()")

	got, err := Parse(data)
	assertNoError(t, err, "Parse()")

	want := validRunbook()

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Parse() = %#v, want %#v", got, want)
	}
}

// TestParseInvalidYAML
func TestParseInvalidYAML(t *testing.T) {
	data, err := os.ReadFile("testdata/invalid_runbook.yaml")
	assertNoError(t, err, "ReadFile()")

	_, err = Parse(data)
	assertErrorContains(t, err, "Parse()", "parse runbook YAML")
}

// TestLoadFile
func TestLoadValidFile(t *testing.T) {
	want := validRunbook()

	got, err := Load("testdata/valid_runbook.yaml")
	assertNoError(t, err, "Load()")

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Load() = %#v, want %#v", got, want)
	}
}

// TestLoadMissingFile
func TestLoadMissingFile(t *testing.T) {
	_, err := Load("testdata/missing.yaml")
	assertErrorContains(t, err, "Load()", "read runbook file")
}

// TestLoadInvalidYAML
func TestLoadInvalidYAML(t *testing.T) {
	_, err := Load("testdata/invalid_runbook.yaml")
	assertErrorContains(t, err, "Load()", "parse runbook YAML")
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
