package tui

import (
	"strings"
	"testing"

	tea "charm.land/bubbletea/v2"

	"github.com/sgrinan/runbeacon/internal/engine"
	"github.com/sgrinan/runbeacon/internal/testutil"
)

func TestUpdateQuitsOnCtrlC(t *testing.T) {
	catalog := engine.Catalog{
		Runbooks: testutil.ValidRunbooks(),
	}

	model := NewModel(catalog)

	msg := tea.KeyPressMsg{
		Code: 'c',
		Mod:  tea.ModCtrl,
	}

	_, cmd := model.Update(msg)

	if cmd == nil {
		t.Fatal("Update() command = nil, want quit command")
	}
}

func TestUpdateSelectsRunbookOnEnter(t *testing.T) {
	catalog := engine.Catalog{
		Runbooks: testutil.ValidRunbooks(),
	}

	model := NewModel(catalog)
	model.list.Select(1)

	msg := tea.KeyPressMsg{
		Code: tea.KeyEnter,
	}

	updatedModel, _ := model.Update(msg)

	updated, ok := updatedModel.(Model)
	if !ok {
		t.Fatalf("Update() model type = %T, want tui.Model", updatedModel)
	}

	if updated.selected == nil {
		t.Fatal("Update() selected = nil, want selected runbook")
	}

	if updated.selected.ID != "kubernetes.pod.pending" {
		t.Fatalf("Update() selected ID = %q, want %q", updated.selected.ID, "kubernetes.pod.pending")
	}
}

func TestUpdateEnterWithNoResults(t *testing.T) {
	catalog := engine.Catalog{
		Runbooks: testutil.ValidRunbooks(),
	}

	model := NewModel(catalog)
	model.list.SetItems(nil)

	msg := tea.KeyPressMsg{
		Code: tea.KeyEnter,
	}

	updatedModel, _ := model.Update(msg)

	updated, ok := updatedModel.(Model)
	if !ok {
		t.Fatalf("Update() model type = %T, want tui.Model", updatedModel)
	}

	if updated.selected != nil {
		t.Fatal("Update() selected != nil, want nil")
	}
}

func TestUpdateReturnsToSearchOnEscape(t *testing.T) {
	runbook := testutil.ValidRunbook()

	model := Model{
		selected: &runbook,
	}

	msg := tea.KeyPressMsg{
		Code: tea.KeyEscape,
	}

	updatedModel, _ := model.Update(msg)

	updated, ok := updatedModel.(Model)
	if !ok {
		t.Fatalf("Update() model type = %T, want tui.Model", updatedModel)
	}

	if updated.selected != nil {
		t.Fatalf("Update() selected = %#v, want nil", updated.selected)
	}
}

func TestUpdateUpdatesComponentSizes(t *testing.T) {
	catalog := engine.Catalog{
		Runbooks: testutil.ValidRunbooks(),
	}

	model := NewModel(catalog)

	updatedModel, _ := model.Update(
		tea.WindowSizeMsg{
			Width:  120,
			Height: 40,
		},
	)

	updated := updatedModel.(Model)

	if updated.viewport.Width() != 120 {
		t.Fatalf("Update() viewport width = %d, want %d", updated.viewport.Width(), 120)
	}

	if updated.viewport.Height() != 39 {
		t.Fatalf("Update() viewport height = %d, want %d", updated.viewport.Height(), 39)
	}

	if updated.list.Width() != 120 {
		t.Fatalf("Update() list width = %d, want %d", updated.list.Width(), 120)
	}

	if updated.list.Height() != 40 {
		t.Fatalf("Update() list height = %d, want %d", updated.list.Height(), 40)
	}
}

func TestUpdateLoadsRunbookIntoViewport(t *testing.T) {
	catalog := engine.Catalog{
		Runbooks: testutil.ValidRunbooks(),
	}

	model := NewModel(catalog)

	updatedModel, _ := model.Update(
		tea.WindowSizeMsg{
			Width:  120,
			Height: 40,
		},
	)
	model = updatedModel.(Model)

	updatedModel, _ = model.Update(
		tea.KeyPressMsg{Code: tea.KeyEnter},
	)
	updated := updatedModel.(Model)

	got := updated.viewport.View()

	if !strings.Contains(got, "Pod in CrashLoopBackOff") {
		t.Fatalf("Update() viewport = %q, want runbook title", got)
	}

	if !strings.Contains(got, "kubectl describe pod") {
		t.Fatalf("Update() viewport = %q, want runbook command", got)
	}
}

func TestUpdateScrollsViewportWhenRunbookIsSelected(t *testing.T) {
	catalog := engine.Catalog{
		Runbooks: testutil.ValidRunbooks(),
	}

	model := NewModel(catalog)

	updatedModel, _ := model.Update(
		tea.WindowSizeMsg{
			Width:  80,
			Height: 5,
		},
	)
	model = updatedModel.(Model)

	updatedModel, _ = model.Update(
		tea.KeyPressMsg{Code: tea.KeyEnter},
	)
	model = updatedModel.(Model)

	before := model.viewport.YOffset()

	updatedModel, _ = model.Update(
		tea.KeyPressMsg{Code: tea.KeyDown},
	)
	updated := updatedModel.(Model)

	if updated.viewport.YOffset() <= before {
		t.Fatalf("Update() viewport offset = %d, want greater than %d", updated.viewport.YOffset(), before)
	}
}
