package tui

import (
	"testing"

	"charm.land/bubbles/v2/list"

	"github.com/sgrinan/runbeacon/internal/testutil"
)

func TestFilterRunbooksMatchesSubstringCaseInsensitive(t *testing.T) {
	targets := []string{
		"Pod in CrashLoopBackOff Diagnose a repeatedly crashing pod",
		"Pod in ImagePullBackOff Diagnose a pod that cannot pull its container image",
		"Pod terminated with OOMKilled Diagnose a container terminated because it exceeded its memory limit",
	}

	got := filterRunbooks("container", targets)
	want := []list.Rank{{Index: 1}, {Index: 2}}

	testutil.DeepEqual(t, got, want, "filterRunbooks()")
}
