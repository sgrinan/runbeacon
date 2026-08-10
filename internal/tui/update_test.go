package tui

import (
	"strings"
	"testing"

	tea "charm.land/bubbletea/v2"

	"github.com/sgrinan/runbeacon/internal/engine"
	"github.com/sgrinan/runbeacon/internal/testutil"
)

func TestUpdateMovesCursorDown(t *testing.T) {
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

func TestUpdateMovesCursorUp(t *testing.T) {
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

func TestUpdateUpdatesSearchQuery(t *testing.T) {
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

func TestUpdateDeletesSearchCharacter(t *testing.T) {
	catalog := engine.Catalog{
		Runbooks: testutil.ValidRunbooks(),
	}

	model := NewModel(catalog)
	model.query = "restart"
	model.results = catalog.Search(model.query)

	msg := tea.KeyPressMsg{
		Code: tea.KeyBackspace,
	}

	updatedModel, _ := model.Update(msg)

	updated, ok := updatedModel.(Model)
	if !ok {
		t.Fatalf("Update() model type = %T, want tui.Model", updatedModel)
	}

	if updated.query != "restar" {
		t.Fatalf("Update() query = %q, want %q", updated.query, "restar")
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

func TestUpdateBackspaceOnEmptyQuery(t *testing.T) {
	catalog := engine.Catalog{
		Runbooks: testutil.ValidRunbooks(),
	}

	model := NewModel(catalog)

	msg := tea.KeyPressMsg{
		Code: tea.KeyBackspace,
	}

	updatedModel, _ := model.Update(msg)

	updated, ok := updatedModel.(Model)
	if !ok {
		t.Fatalf("Update() model type = %T, want tui.Model", updatedModel)
	}

	if updated.query != "" {
		t.Fatalf("Update() query = %q, want %q", updated.query, "")
	}
}

func TestUpdateCursorStaysAtTop(t *testing.T) {
	catalog := engine.Catalog{
		Runbooks: testutil.ValidRunbooks(),
	}

	model := NewModel(catalog)

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

func TestUpdateCursorStaysAtBottom(t *testing.T) {
	catalog := engine.Catalog{
		Runbooks: testutil.ValidRunbooks(),
	}

	model := NewModel(catalog)
	model.cursor = len(model.results) - 1

	msg := tea.KeyPressMsg{
		Code: tea.KeyDown,
	}

	updatedModel, _ := model.Update(msg)

	updated, ok := updatedModel.(Model)
	if !ok {
		t.Fatalf("Update() model type = %T, want tui.Model", updatedModel)
	}

	want := len(model.results) - 1

	if updated.cursor != want {
		t.Fatalf("Update() cursor = %d, want %d", updated.cursor, want)
	}
}

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
	model.cursor = 1

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
	model.query = "does-not-exist"
	model.results = catalog.Search(model.query)

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

	if len(model.results) != 0 {
		t.Fatalf("test setup: result count = %d, want 0", len(model.results))
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

func TestUpdateUpdatesViewportSize(t *testing.T) {
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

	if updated.viewport.Height() != 40 {
		t.Fatalf("Update() viewport height = %d, want %d", updated.viewport.Height(), 40)
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
