package ui

import (
	"context"
	"time"

	apitypes "github.com/appclacks/go-types"
	tea "github.com/charmbracelet/bubbletea"
)

type listHealthcheckMsg struct {
	failures     map[string]int
	healthchecks []apitypes.Healthcheck
}

func (m *model) listHealthchecks() tea.Msg {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()
	result, err := m.client.ListHealthchecks(ctx)
	if err != nil {
		return errMsg{err}
	}

	start := time.Now().UTC().Add(-10 * time.Minute)
	end := time.Now().UTC()
	failures := make(map[string]int)
	f := false
	for _, c := range result.Result {
		resultCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		results, err := m.client.ListHealthchecksResults(resultCtx, apitypes.ListHealthchecksResultsInput{
			StartDate:     start,
			EndDate:       end,
			HealthcheckID: c.ID,
			Success:       &f,
		})
		if err != nil {
			return errMsg{err}
		}
		failures[c.ID] = len(results.Result)
	}
	return listHealthcheckMsg{
		failures:     failures,
		healthchecks: result.Result,
	}
}
