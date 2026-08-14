package tui

import (
	"fmt"
	"strings"

	"github.com/sgrinan/runbeacon/internal/models"
)

func renderRunbook(runbook models.Runbook) string {
	var builder strings.Builder

	fmt.Fprintf(&builder, "%s\n", titleStyle.Render(runbook.Title))
	fmt.Fprintf(&builder, "%s\n\n", mutedStyle.Render(runbook.ID))
	fmt.Fprintf(&builder, "%s\n\n", runbook.Description)

	fmt.Fprintf(&builder, "%s\n", sectionStyle.Render("Symptoms"))

	for _, symptom := range runbook.Symptoms {
		fmt.Fprintf(&builder, "- %s\n", symptom.Name)

		keywords := "  Keywords: " + strings.Join(symptom.Keywords, ", ")
		fmt.Fprintf(&builder, "%s\n\n", mutedStyle.Render(keywords))
	}

	fmt.Fprintf(&builder, "%s\n", sectionStyle.Render("Steps"))

	for i, step := range runbook.Steps {
		stepTitle := fmt.Sprintf("%d. %s", i+1, step.Title)

		fmt.Fprintf(&builder, "%s\n", stepTitleStyle.Render(stepTitle))
		fmt.Fprintf(&builder, "   %s\n", mutedStyle.Render(step.Description))

		if step.Command != "" {
			command := "   $ " + step.Command
			fmt.Fprintf(&builder, "%s\n", commandStyle.Render(command))
		}

		builder.WriteString("\n")
	}

	return builder.String()
}
