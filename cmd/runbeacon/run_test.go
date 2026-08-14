package main

import (
	"bytes"
	"testing"

	"github.com/sgrinan/runbeacon/internal/engine"
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
			name:       "unknown command",
			args:       []string{"banana"},
			wantCode:   2,
			wantStderr: "unknown command: banana\n",
		},
		{
			name:       "search without query",
			args:       []string{"search"},
			wantCode:   2,
			wantStderr: "usage: runbeacon search <query>\n",
		},
		{
			name:       "valid search",
			args:       []string{"search", "crash"},
			wantCode:   0,
			wantStdout: "kubernetes.pod.crashloopbackoff\tPod in CrashLoopBackOff\n",
		},
	}

	catalog := engine.Catalog{
		Runbooks: testutil.Runbooks(),
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var stdout bytes.Buffer
			var stderr bytes.Buffer

			gotCode := run(tt.args, &stdout, &stderr, catalog)

			testutil.Equal(t, gotCode, tt.wantCode, "run() code")
			testutil.EqualString(t, stdout.String(), tt.wantStdout, "run() stdout")
			testutil.EqualString(t, stderr.String(), tt.wantStderr, "run() stderr")
		})
	}
}
