package testutil

import "github.com/sgrinan/runbeacon/internal/models"

func ValidRunbook() models.Runbook {
	return models.Runbook{
		ID:          "kubernetes.pod.crashloopbackoff",
		Title:       "Pod in CrashLoopBackOff",
		Description: "Diagnose a repeatedly crashing pod.",
		Symptoms: []models.Symptom{
			{
				ID:       "kubernetes.pod.crash_loop",
				Name:     "CrashLoopBackOff",
				Keywords: []string{"crash", "restart", "backoff"},
			},
		},
		Steps: []models.Step{
			{
				Title:       "Describe pod",
				Description: "Inspect pod events and container state.",
				Command:     "kubectl describe pod",
			},
		},
	}
}
