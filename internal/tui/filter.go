package tui

import (
	"strings"

	"charm.land/bubbles/v2/list"
)

func filterRunbooks(term string, targets []string) []list.Rank {
	term = strings.ToLower(strings.TrimSpace(term))

	var ranks []list.Rank

	for i, target := range targets {
		if strings.Contains(strings.ToLower(target), term) {
			ranks = append(ranks, list.Rank{Index: i})
		}
	}

	return ranks
}
