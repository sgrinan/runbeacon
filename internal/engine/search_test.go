package engine

import (
	"testing"

	"github.com/sgrinan/runbeacon/internal/models"
	"github.com/sgrinan/runbeacon/internal/testutil"
)

func TestSearch(t *testing.T) {
	tests := []struct {
		name    string
		query   string
		wantIDs []string
	}{
		{
			name:    "matches id",
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
			name:    "matches symptoms id",
			query:   "crash_loop",
			wantIDs: []string{"kubernetes.pod.crashloopbackoff"},
		},
		{
			name:    "matches symptoms name",
			query:   "CrashLoopBackOff",
			wantIDs: []string{"kubernetes.pod.crashloopbackoff"},
		},
		{
			name:    "matches symptoms keyword",
			query:   "restart",
			wantIDs: []string{"kubernetes.pod.crashloopbackoff"},
		},
		{
			name:    "matches steps title",
			query:   "Describe pod",
			wantIDs: []string{"kubernetes.pod.crashloopbackoff"},
		},
		{
			name:    "matches steps description",
			query:   "Inspect pod",
			wantIDs: []string{"kubernetes.pod.crashloopbackoff"},
		},
		{
			name:    "matches steps command",
			query:   "kubectl",
			wantIDs: []string{"kubernetes.pod.crashloopbackoff"},
		},
	}

	catalog := Catalog{
		Runbooks: []models.Runbook{
			testutil.ValidRunbook(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := catalog.Search(tt.query)

			if len(got) != len(tt.wantIDs) {
				t.Fatalf("Search(%q) result count = %d, want %d", tt.query, len(got), len(tt.wantIDs))
			}

			if got[0].ID != tt.wantIDs[0] {
				t.Fatalf("Search(%q) ID = %q, want %q", tt.query, got[0].ID, tt.wantIDs[0])
			}
		})
	}
}
