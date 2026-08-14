package render

import (
	"testing"

	"github.com/sgrinan/runbeacon/internal/testutil"
)

func TestRunbook(t *testing.T) {
	tests := []struct {
		name    string
		command string
		want    string
	}{
		{
			name:    "with command",
			command: "kubectl describe pod",
			want: `Pod in CrashLoopBackOff
kubernetes.pod.crashloopbackoff

Diagnose a repeatedly crashing pod.

Symptoms:
- CrashLoopBackOff
  Keywords: crash, restart, backoff

Steps:
1. Describe pod
   Inspect pod events and container state.
   $ kubectl describe pod

`,
		},
		{
			name: "without command",
			want: `Pod in CrashLoopBackOff
kubernetes.pod.crashloopbackoff

Diagnose a repeatedly crashing pod.

Symptoms:
- CrashLoopBackOff
  Keywords: crash, restart, backoff

Steps:
1. Describe pod
   Inspect pod events and container state.

`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runbook := testutil.CrashLoopBackOffRunbook()
			runbook.Steps[0].Command = tt.command

			got := Runbook(runbook)

			testutil.EqualString(t, got, tt.want, "Runbook()")
		})
	}
}
