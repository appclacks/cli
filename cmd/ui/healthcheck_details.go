package ui

import (
	"context"
	"fmt"
	"strings"
	"time"

	apitypes "github.com/appclacks/go-types"
	tea "github.com/charmbracelet/bubbletea"
)

type healthcheckDetailsMsg struct {
	failures    []apitypes.HealthcheckResult
	healthcheck apitypes.Healthcheck
}

func detailsFailures(failures []apitypes.HealthcheckResult) string {

	failure := failures[len(failures)-1]
	f := ""
	f += fmt.Sprintf(" Date: %s\n", failure.CreatedAt)
	if len(failure.Labels) != 0 {
		var labels []string
		for k, v := range failure.Labels {
			labels = append(labels, fmt.Sprintf("%s=%s", k, v))
		}
		f += fmt.Sprintf(" Labels: %s\n\n", strings.Join(labels, ", "))
	}

	f += fmt.Sprintf(" Message: %s\n\n", failure.Message)
	return f

}

func detailString(healthcheck apitypes.Healthcheck, failures []apitypes.HealthcheckResult) string {
	result := fmt.Sprintf(" Name: %s\n", healthcheck.Name)
	result += fmt.Sprintf(" ID: %s\n", healthcheck.ID)
	result += fmt.Sprintf(" Type: %s\n", healthcheck.Type)
	if healthcheck.Description != "" {
		result += fmt.Sprintf(" Description: %s\n", healthcheck.Description)
	}
	result += fmt.Sprintf(" Interval: %s\n", healthcheck.Interval)
	result += fmt.Sprintf(" Timeout: %s\n", healthcheck.Timeout)
	result += fmt.Sprintf(" Created At: %s\n", healthcheck.CreatedAt)
	result += fmt.Sprintf(" Enabled: %t\n", healthcheck.Enabled)
	if len(healthcheck.Labels) != 0 {
		var labels []string
		for k, v := range healthcheck.Labels {
			labels = append(labels, fmt.Sprintf("%s=%s", k, v))
		}
		result += fmt.Sprintf(" Labels: %s\n", strings.Join(labels, ", "))
	}
	result += fmt.Sprintf("\n Number of Failures (last 10 min): %d\n", len(failures))
	return result
}

func (m *model) getHealthchecksDetails() tea.Msg {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	healthcheck, err := m.client.GetHealthcheck(ctx, apitypes.GetHealthcheckInput{
		ID: m.selectedHealthcheckID,
	})
	if err != nil {
		return errMsg{err}
	}

	resultCtx, resultCancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer resultCancel()

	start := time.Now().UTC().Add(-10 * time.Minute)
	end := time.Now().UTC()
	f := false
	results, err := m.client.ListHealthchecksResults(resultCtx, apitypes.ListHealthchecksResultsInput{
		StartDate:     start,
		EndDate:       end,
		HealthcheckID: m.selectedHealthcheckID,
		Success:       &f,
	})
	if err != nil {
		return errMsg{err}
	}

	return healthcheckDetailsMsg{
		healthcheck: healthcheck,
		failures:    results.Result,
	}
}
