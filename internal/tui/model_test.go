package tui

import (
	"reflect"
	"testing"

	"github.com/sgrinan/runbeacon/internal/engine"
	"github.com/sgrinan/runbeacon/internal/testutil"
)

func TestNewModel(t *testing.T) {
	catalog := engine.Catalog{
		Runbooks: testutil.ValidRunbooks(),
	}

	model := NewModel(catalog)

	if !reflect.DeepEqual(model.catalog, catalog) {
		t.Fatalf("NewModel() catalog = %#v, want %#v", model.catalog, catalog)
	}

	if model.query != "" {
		t.Fatalf("NewModel() query = %q, want %q", model.query, "")
	}

	if !reflect.DeepEqual(model.results, catalog.Runbooks) {
		t.Fatalf("NewModel() results = %#v, want %#v", model.results, catalog.Runbooks)
	}

	if model.cursor != 0 {
		t.Fatalf("NewModel() cursor = %d, want %d", model.cursor, 0)
	}

	if model.selected != nil {
		t.Fatalf("NewModel() selected = %#v, want nil", model.selected)
	}
}
