package parser

import (
	"os"
	"testing"

	"github.com/sgrinan/runbeacon/internal/testutil"
)

func TestParseValidYAML(t *testing.T) {
	data, err := os.ReadFile("testdata/validYAML.yaml")
	testutil.NoError(t, err, "ReadFile()")

	got, err := Parse(data)
	testutil.NoError(t, err, "Parse()")

	want := testutil.CrashLoopBackOffRunbook()

	testutil.DeepEqual(t, got, want, "Parse()")
}

func TestParseInvalidYAML(t *testing.T) {
	data, err := os.ReadFile("testdata/invalidYAML.yaml")
	testutil.NoError(t, err, "ReadFile()")

	_, err = Parse(data)
	testutil.ErrorContains(t, err, "Parse()", "parse runbook YAML")
}

func TestLoadValidFile(t *testing.T) {
	got, err := Load("testdata/validYAML.yaml")
	testutil.NoError(t, err, "Load()")

	want := testutil.CrashLoopBackOffRunbook()

	testutil.DeepEqual(t, got, want, "Load()")
}

func TestLoadMissingFile(t *testing.T) {
	_, err := Load("testdata/missing.yaml")
	testutil.ErrorContains(t, err, "Load()", "read runbook file")
}

func TestLoadInvalidYAML(t *testing.T) {
	_, err := Load("testdata/invalidYAML.yaml")
	testutil.ErrorContains(t, err, "Load()", "parse runbook YAML")
}
