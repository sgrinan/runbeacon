package render

import (
	"fmt"
	"strings"

	"github.com/sgrinan/runbeacon/internal/models"
)

func Runbook(runbook models.Runbook) string {
	var builder strings.Builder
	fmt.Fprintf(&builder, "%s\n", runbook.Title)
	fmt.Fprintf(&builder, "%s\n\n", runbook.ID)
	fmt.Fprintf(&builder, "%s\n\n", runbook.Description)

	builder.WriteString("Symptoms:\n")

	for _, symptom := range runbook.Symptoms {
		fmt.Fprintf(&builder, "- %s\n", symptom.Name)
		fmt.Fprintf(&builder, "  Keywords: ")
		fmt.Fprintf(&builder, "%s\n\n", strings.Join(symptom.Keywords, ", "))
	}

	builder.WriteString("Steps:\n")

	for i, step := range runbook.Steps {
		fmt.Fprintf(&builder, "%d. %s\n", i+1, step.Title)
		fmt.Fprintf(&builder, "   %s\n", step.Description)
		if step.Command != "" {
			fmt.Fprintf(&builder, "   $ %s\n", step.Command)
		}
	}

	return builder.String()
}
