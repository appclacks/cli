package types

import (
	"time"
)

type ListHealthchecksResultsInput struct {
	StartDate     time.Time `query:"start-date" validate:"required"`
	EndDate       time.Time `query:"end-date" validate:"required"`
	HealthcheckID string    `query:"healthcheck-id" validate:"omitempty,uuid"`
	Page          int       `query:"page" validate:"omitempty,min=1"`
	Success       *bool     `query:"success"`
}

type HealthcheckResult struct {
	ID            string            `json:"id"`
	Success       bool              `json:"success"`
	Labels        map[string]string `json:"labels,omitempty"`
	CreatedAt     time.Time         `json:"created-at"`
	Summary       string            `json:"summary"`
	Message       string            `json:"message"`
	HealthcheckID string            `json:"healthcheck-id"`
}

type ListHealthchecksResultsOutput struct {
	Result []HealthcheckResult `json:"result"`
}
