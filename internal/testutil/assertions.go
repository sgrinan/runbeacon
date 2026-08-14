package testutil

import (
	"reflect"
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

func Equal[T comparable](t *testing.T, got, want T, operation string) {
	t.Helper()

	if got != want {
		t.Fatalf("%s = %v, want %v", operation, got, want)
	}
}

func EqualString(t *testing.T, got, want, operation string) {
	t.Helper()

	if got != want {
		t.Fatalf("%s = %q, want %q", operation, got, want)
	}
}

func DeepEqual(t *testing.T, got, want any, operation string) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("%s = %#v, want %#v", operation, got, want)
	}
}

func StringContains(t *testing.T, got, want, operation string) {
	t.Helper()

	if !strings.Contains(got, want) {
		t.Fatalf("%s = %q, want it to contain %q", operation, got, want)
	}
}
