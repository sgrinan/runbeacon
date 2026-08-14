package tui

import (
	"strings"
	"testing"

	"github.com/sgrinan/runbeacon/internal/testutil"
)

func TestRenderRunbook(t *testing.T) {
	runbook := testutil.CrashLoopBackOffRunbook()

	got := renderRunbook(runbook)

	wantTerms := []string{
		"Pod in CrashLoopBackOff",
		"kubernetes.pod.crashloopbackoff",
		"Diagnose a repeatedly crashing pod.",
		"CrashLoopBackOff",
		"crash, restart, backoff",
		"Describe pod",
		"kubectl describe pod",
	}

	for _, term := range wantTerms {
		testutil.StringContains(t, got, term, "renderRunbook()")
	}
}

func TestRenderRunbookOmitsEmptyCommand(t *testing.T) {
	runbook := testutil.CrashLoopBackOffRunbook()
	runbook.Steps[0].Command = ""

	got := renderRunbook(runbook)

	if strings.Contains(got, "$ kubectl describe pod") {
		t.Fatalf("renderRunbook() = %q, want command omitted", got)
	}
}
