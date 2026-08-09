package testutil

import (
	"github.com/sgrinan/runbeacon/internal/engine"
	"github.com/sgrinan/runbeacon/internal/models"
)

func ValidCatalog() engine.Catalog {
	return engine.Catalog{
		Runbooks: []models.Runbook{
			ValidRunbook(),
			{
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
			},
		},
	}
}
