package render

import (
	"testing"

	"github.com/sgrinan/runbeacon/internal/testutil"
)

func TestRunbook(t *testing.T) {
	runbook := testutil.ValidRunbook()

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

	got := Runbook(runbook)

	if got != want {
		t.Fatalf("Runbook() = %q, want %q", got, want)
	}
}

func TestRunbookOmitsEmptyCommand(t *testing.T) {
	runbook := testutil.ValidRunbook()
	runbook.Steps[0].Command = ""

	want := `Pod in CrashLoopBackOff
kubernetes.pod.crashloopbackoff

Diagnose a repeatedly crashing pod.

Symptoms:
- CrashLoopBackOff
  Keywords: crash, restart, backoff

Steps:
1. Describe pod
   Inspect pod events and container state.

`

	got := Runbook(runbook)

	if got != want {
		t.Fatalf("Runbook() = %q, want %q", got, want)
	}
}
