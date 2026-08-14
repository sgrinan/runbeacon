package tui

import "charm.land/lipgloss/v2"

var (
	appStyle = lipgloss.NewStyle().
			Padding(1, 2)

	stepTitleStyle = lipgloss.NewStyle().
			Bold(true)

	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("205"))

	sectionStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("205"))

	mutedStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("240"))

	commandStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("86"))

	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("240"))
)
