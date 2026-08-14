package engine

import (
	"testing"

	"github.com/sgrinan/runbeacon/internal/testutil"
)

func TestSearch(t *testing.T) {
	tests := []struct {
		name    string
		query   string
		wantIDs []string
	}{
		{
			name:    "matches ID",
			query:   "pod.crashloop",
			wantIDs: []string{"kubernetes.pod.crashloopbackoff"},
		},
		{
			name:    "matches title",
			query:   "Pod in Crash",
			wantIDs: []string{"kubernetes.pod.crashloopbackoff"},
		},
		{
			name:    "matches description",
			query:   "crashing pod",
			wantIDs: []string{"kubernetes.pod.crashloopbackoff"},
		},
		{
			name:    "matches symptom ID",
			query:   "crash_loop",
			wantIDs: []string{"kubernetes.pod.crashloopbackoff"},
		},
		{
			name:    "matches symptom name",
			query:   "CrashLoopBackOff",
			wantIDs: []string{"kubernetes.pod.crashloopbackoff"},
		},
		{
			name:    "matches symptom keyword",
			query:   "restart",
			wantIDs: []string{"kubernetes.pod.crashloopbackoff"},
		},
		{
			name:    "ignores case",
			query:   "CRASHLOOPBACKOFF",
			wantIDs: []string{"kubernetes.pod.crashloopbackoff"},
		},
		{
			name:    "trims query spaces",
			query:   "   restart   ",
			wantIDs: []string{"kubernetes.pod.crashloopbackoff"},
		},
		{
			name:    "no match",
			query:   "network-timeout",
			wantIDs: []string{},
		},
		{
			name:  "empty query",
			query: "",
			wantIDs: []string{
				"kubernetes.pod.crashloopbackoff",
				"kubernetes.pod.pending",
			},
		},
		{
			name:  "matches multiple runbooks",
			query: "pod",
			wantIDs: []string{
				"kubernetes.pod.crashloopbackoff",
				"kubernetes.pod.pending",
			},
		},
	}

	catalog := Catalog{
		Runbooks: testutil.Runbooks(),
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := catalog.Search(tt.query)

			gotIDs := make([]string, 0, len(got))
			for _, runbook := range got {
				gotIDs = append(gotIDs, runbook.ID)
			}

			testutil.DeepEqual(t, gotIDs, tt.wantIDs, "Search() result IDs")
		})
	}
}
