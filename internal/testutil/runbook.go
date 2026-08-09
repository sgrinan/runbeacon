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

func ValidPendingRunbook() models.Runbook {
	return models.Runbook{
		ID:          "kubernetes.pod.pending",
		Title:       "Pod Pending",
		Description: "Diagnose a pod stuck in Pending state.",
		Symptoms: []models.Symptom{
			{
				ID:       "kubernetes.pod.pending",
				Name:     "Pending",
				Keywords: []string{"pod", "pending", "scheduling"},
			},
		},
		Steps: []models.Step{
			{
				Title:       "Describe pod",
				Description: "Inspect scheduling events.",
				Command:     "kubectl describe pod",
			},
		},
	}
}

func ValidRunbooks() []models.Runbook {
	return []models.Runbook{
		ValidRunbook(),
		ValidPendingRunbook(),
	}
}
