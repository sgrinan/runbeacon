package tui

import "charm.land/lipgloss/v2"

var (
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
