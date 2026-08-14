package tui

import (
	"strings"
	"testing"

	"github.com/sgrinan/runbeacon/internal/testutil"
)

func TestItemFilterValueIncludesSearchableFields(t *testing.T) {
	item := runbookItem{runbook: testutil.ValidRunbook()}

	got := item.FilterValue()

	wantTerms := []string{
		"Pod in CrashLoopBackOff",
		"kubernetes.pod.crashloopbackoff",
		"CrashLoopBackOff",
		"crash",
		"restart",
		"backoff",
		"Diagnose a repeatedly crashing pod.",
	}

	for _, term := range wantTerms {
		if !strings.Contains(got, term) {
			t.Fatalf("FilterValue() = %q, want to contain %q", got, term)
		}
	}
}
