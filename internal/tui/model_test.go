package tui

import (
	"reflect"
	"testing"

	"charm.land/bubbletea/v2"
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
}

func TestModelMovesCursorDown(t *testing.T) {
	catalog := engine.Catalog{
		Runbooks: testutil.ValidRunbooks(),
	}

	model := NewModel(catalog)

	msg := tea.KeyPressMsg{
		Code: tea.KeyDown,
	}

	updatedModel, _ := model.Update(msg)

	updated, ok := updatedModel.(Model)
	if !ok {
		t.Fatalf("Update() model type = %T, want tui.Model", updatedModel)
	}

	if updated.cursor != 1 {
		t.Fatalf("Update() cursor = %d, want 1", updated.cursor)
	}
}

func TestModelMovesCursorUp(t *testing.T) {
	catalog := engine.Catalog{
		Runbooks: testutil.ValidRunbooks(),
	}

	model := NewModel(catalog)
	model.cursor = 1

	msg := tea.KeyPressMsg{
		Code: tea.KeyUp,
	}

	updatedModel, _ := model.Update(msg)

	updated, ok := updatedModel.(Model)
	if !ok {
		t.Fatalf("Update() model type = %T, want tui.Model", updatedModel)
	}

	if updated.cursor != 0 {
		t.Fatalf("Update() cursor = %d, want 0", updated.cursor)
	}
}

func TestModelUpdatesSearchQuery(t *testing.T) {
	catalog := engine.Catalog{
		Runbooks: testutil.ValidRunbooks(),
	}

	model := NewModel(catalog)

	msg := tea.KeyPressMsg{
		Text: "restart",
	}

	updatedModel, _ := model.Update(msg)

	updated, ok := updatedModel.(Model)
	if !ok {
		t.Fatalf("Update() model type = %T, want tui.Model", updatedModel)
	}

	if updated.query != "restart" {
		t.Fatalf("Update() query = %q, want %q", updated.query, "restart")
	}

	if len(updated.results) != 1 {
		t.Fatalf("Update() result count = %d, want 1", len(updated.results))
	}

	if updated.results[0].ID != "kubernetes.pod.crashloopbackoff" {
		t.Fatalf("Update() result ID = %q, want %q", updated.results[0].ID, "kubernetes.pod.crashloopbackoff")
	}

	if updated.cursor != 0 {
		t.Fatalf("Update() cursor = %d, want 0", updated.cursor)
	}
}
