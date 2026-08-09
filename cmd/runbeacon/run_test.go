package main

import (
	"bytes"
	"testing"

	"github.com/sgrinan/runbeacon/internal/engine"
	"github.com/sgrinan/runbeacon/internal/models"
	"github.com/sgrinan/runbeacon/internal/testutil"
)

func TestRun(t *testing.T) {
	tests := []struct {
		name       string
		args       []string
		wantCode   int
		wantStdout string
		wantStderr string
	}{
		{
			name:       "no arguments",
			args:       []string{},
			wantCode:   2,
			wantStdout: "",
			wantStderr: "usage: runbeacon <command>\n",
		},
		{
			name:       "unknown command",
			args:       []string{"banana"},
			wantCode:   2,
			wantStdout: "",
			wantStderr: "unknown command: banana\n",
		},
		{
			name:       "search without query",
			args:       []string{"search"},
			wantCode:   2,
			wantStdout: "",
			wantStderr: "usage: runbeacon search <query>\n",
		},
		{
			name:       "valid search",
			args:       []string{"search", "crash"},
			wantCode:   0,
			wantStdout: "kubernetes.pod.crashloopbackoff\tPod in CrashLoopBackOff\n",
			wantStderr: "",
		},
	}

	catalog := engine.Catalog{
		Runbooks: []models.Runbook{
			testutil.ValidRunbook(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var stdout bytes.Buffer
			var stderr bytes.Buffer

			gotCode := run(tt.args, &stdout, &stderr, catalog)
			if gotCode != tt.wantCode {
				t.Fatalf("run() code = %d, want %d", gotCode, tt.wantCode)
			}

			if stdout.String() != tt.wantStdout {
				t.Fatalf("run() stdout = %q, want %q", stdout.String(), tt.wantStdout)
			}

			if stderr.String() != tt.wantStderr {
				t.Fatalf("run() stderr = %q, want %q", stderr.String(), tt.wantStderr)
			}
		})
	}
}
