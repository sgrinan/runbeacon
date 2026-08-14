package tui

import (
	"testing"

	"charm.land/bubbles/v2/list"
)

func TestFilterRunbooksMatchesSubstringCaseInsensitive(t *testing.T) {
	targets := []string{
		"Pod in CrashLoopBackOff Diagnose a repeatedly crashing pod",
		"Pod in ImagePullBackOff Diagnose a pod that cannot pull its container image",
		"Pod terminated with OOMKilled Diagnose a container terminated because it exceeded its memory limit",
	}

	got := filterRunbooks("container", targets)

	want := []list.Rank{{Index: 1}, {Index: 2}}

	if len(got) != len(want) {
		t.Fatalf("filterRunbooks() result count = %d, want %d", len(got), len(want))
	}

	for i := range want {
		if got[i].Index != want[i].Index {
			t.Fatalf("filterRunbooks() result[%d].Index = %d, want %d", i, got[i].Index, want[i].Index)
		}
	}
}
