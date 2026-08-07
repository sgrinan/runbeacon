package testutil

import (
	"strings"
	"testing"
)

func NoError(t *testing.T, err error, operation string) {
	t.Helper()

	if err != nil {
		t.Fatalf("%s: %v", operation, err)
	}
}

func ErrorContains(t *testing.T, err error, operation string, want string) {
	t.Helper()

	if err == nil {
		t.Fatalf("%s: want an error", operation)
	}

	if !strings.Contains(err.Error(), want) {
		t.Fatalf("%s: error = %q, want it to contain %q", operation, err, want)
	}
}
