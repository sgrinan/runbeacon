package tui

import (
	"testing"

	tea "charm.land/bubbletea/v2"

	"github.com/sgrinan/runbeacon/internal/engine"
	"github.com/sgrinan/runbeacon/internal/testutil"
)

func requireModel(t *testing.T, model tea.Model) Model {
	t.Helper()

	updated, ok := model.(Model)
	if !ok {
		t.Fatalf("model type = %T, want tui.Model", model)
	}

	return updated
}

func TestUpdateQuitsOnCtrlC(t *testing.T) {
	catalog := engine.Catalog{
		Runbooks: testutil.Runbooks(),
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
		Runbooks: testutil.Runbooks(),
	}

	model := NewModel(catalog)
	model.list.Select(1)

	updatedModel, _ := model.Update(
		tea.KeyPressMsg{Code: tea.KeyEnter},
	)
	updated := requireModel(t, updatedModel)

	if updated.selected == nil {
		t.Fatal("Update() selected = nil, want selected runbook")
	}

	testutil.Equal(t, updated.selected.ID, "kubernetes.pod.pending", "Update() selected ID")
}

func TestUpdateEnterWithNoResults(t *testing.T) {
	catalog := engine.Catalog{
		Runbooks: testutil.Runbooks(),
	}

	model := NewModel(catalog)
	model.list.SetItems(nil)

	updatedModel, _ := model.Update(
		tea.KeyPressMsg{Code: tea.KeyEnter},
	)
	updated := requireModel(t, updatedModel)

	if updated.selected != nil {
		t.Fatal("Update() selected != nil, want nil")
	}
}

func TestUpdateReturnsToSearchOnEscape(t *testing.T) {
	runbook := testutil.CrashLoopBackOffRunbook()

	model := Model{
		selected: &runbook,
	}

	updatedModel, _ := model.Update(
		tea.KeyPressMsg{Code: tea.KeyEscape},
	)
	updated := requireModel(t, updatedModel)

	if updated.selected != nil {
		t.Fatalf("Update() selected = %#v, want nil", updated.selected)
	}
}

func TestUpdateUpdatesComponentSizes(t *testing.T) {
	catalog := engine.Catalog{
		Runbooks: testutil.Runbooks(),
	}

	model := NewModel(catalog)

	updatedModel, _ := model.Update(
		tea.WindowSizeMsg{
			Width:  120,
			Height: 40,
		},
	)
	updated := requireModel(t, updatedModel)

	testutil.Equal(t, updated.viewport.Width(), 116, "Update() viewport width")
	testutil.Equal(t, updated.viewport.Height(), 37, "Update() viewport height")
	testutil.Equal(t, updated.list.Width(), 116, "Update() list width")
	testutil.Equal(t, updated.list.Height(), 38, "Update() list height")
}

func TestUpdateLoadsRunbookIntoViewport(t *testing.T) {
	catalog := engine.Catalog{
		Runbooks: testutil.Runbooks(),
	}

	model := NewModel(catalog)

	updatedModel, _ := model.Update(
		tea.WindowSizeMsg{
			Width:  120,
			Height: 40,
		},
	)
	model = requireModel(t, updatedModel)

	updatedModel, _ = model.Update(
		tea.KeyPressMsg{Code: tea.KeyEnter},
	)
	updated := requireModel(t, updatedModel)

	got := updated.viewport.View()

	testutil.StringContains(t, got, "Pod in CrashLoopBackOff", "Update() viewport")
	testutil.StringContains(t, got, "kubectl describe pod", "Update() viewport")
}

func TestUpdateScrollsViewportWhenRunbookIsSelected(t *testing.T) {
	catalog := engine.Catalog{
		Runbooks: testutil.Runbooks(),
	}

	model := NewModel(catalog)

	updatedModel, _ := model.Update(
		tea.WindowSizeMsg{
			Width:  80,
			Height: 5,
		},
	)
	model = requireModel(t, updatedModel)

	updatedModel, _ = model.Update(
		tea.KeyPressMsg{Code: tea.KeyEnter},
	)
	model = requireModel(t, updatedModel)

	before := model.viewport.YOffset()

	updatedModel, _ = model.Update(
		tea.KeyPressMsg{Code: tea.KeyDown},
	)
	updated := requireModel(t, updatedModel)

	if updated.viewport.YOffset() <= before {
		t.Fatalf("Update() viewport offset = %d, want greater than %d", updated.viewport.YOffset(), before)
	}
}
