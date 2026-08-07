package parser

import (
	"os"
	"reflect"
	"testing"

	"github.com/sgrinan/runbeacon/internal/testutil"
)

func TestParseValidYAML(t *testing.T) {
	data, err := os.ReadFile("testdata/valid_runbook.yaml")
	testutil.NoError(t, err, "ReadFile()")

	got, err := Parse(data)
	testutil.NoError(t, err, "Parse()")

	want := testutil.ValidRunbook()

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Parse() = %#v, want %#v", got, want)
	}
}

func TestParseInvalidYAML(t *testing.T) {
	data, err := os.ReadFile("testdata/invalid_runbook.yaml")
	testutil.NoError(t, err, "ReadFile()")

	_, err = Parse(data)
	testutil.ErrorContains(t, err, "Parse()", "parse runbook YAML")
}

func TestLoadValidFile(t *testing.T) {
	want := testutil.ValidRunbook()

	got, err := Load("testdata/valid_runbook.yaml")
	testutil.NoError(t, err, "Load()")

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Load() = %#v, want %#v", got, want)
	}
}

func TestLoadMissingFile(t *testing.T) {
	_, err := Load("testdata/missing.yaml")
	testutil.ErrorContains(t, err, "Load()", "read runbook file")
}

func TestLoadInvalidYAML(t *testing.T) {
	_, err := Load("testdata/invalid_runbook.yaml")
	testutil.ErrorContains(t, err, "Load()", "parse runbook YAML")
}
