package tui

import (
	"testing"

	"github.com/sgrinan/runbeacon/internal/testutil"
)

func TestRunbookView(t *testing.T) {
	runbook := testutil.CrashLoopBackOffRunbook()

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

	testutil.EqualString(t, got, want, "runbookView()")
}
