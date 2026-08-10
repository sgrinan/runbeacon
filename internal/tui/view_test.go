package tui

import (
	"testing"

	"github.com/sgrinan/runbeacon/internal/engine"
	"github.com/sgrinan/runbeacon/internal/testutil"
)

func TestSearchView(t *testing.T) {
	catalog := engine.Catalog{
		Runbooks: testutil.ValidRunbooks(),
	}

	model := NewModel(catalog)
	model.query = "pod"
	model.results = catalog.Search(model.query)

	want := `Search: pod

>  Pod in CrashLoopBackOff
   Pod Pending
`

	got := model.searchView()

	if got != want {
		t.Fatalf("searchView = %q, want %q", got, want)
	}
}

func TestRunbookView(t *testing.T) {
	runbook := testutil.ValidRunbook()

	model := Model{
		selected: &runbook,
	}

	want := `Pod in CrashLoopBackOff
kubernetes.pod.crashloopbackoff

Diagnose a repeatedly crashing pod.

Symptoms:
- CrashLoopBackOff
  Keywords: crash, restart, backoff

Steps:
1. Describe pod
   Inspect pod events and container state.
   $ kubectl describe pod
`

	got := model.runbookView()

	if got != want {
		t.Fatalf("runbookView() = %q, want %q", got, want)
	}
}
