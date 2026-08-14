package tui

import (
	"testing"

	"github.com/sgrinan/runbeacon/internal/testutil"
)

func TestItemFilterValueIncludesSearchableFields(t *testing.T) {
	item := runbookItem{runbook: testutil.CrashLoopBackOffRunbook()}

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
		testutil.StringContains(t, got, term, "FilterValue()")
	}
}
