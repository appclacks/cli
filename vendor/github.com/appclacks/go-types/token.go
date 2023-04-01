package types

import (
	"time"
)

type Permissions struct {
	Actions []string `json:"actions,omitempty"`
}

type APIToken struct {
	ID          string      `json:"id"`
	Name        string      `json:"name"`
	Token       string      `json:"token"`
	Description string      `json:"description"`
	Permissions Permissions `json:"permissions"`
	CreatedAt   time.Time   `json:"created-at"`
	ExpiresAt   time.Time   `json:"expires-at"`
	TTL         string      `json:"ttl"`
}

type CreateAPITokenInput struct {
	Name        string      `json:"name" description:"Token name" validate:"required,max=255"`
	Description string      `json:"description" description:"Token description" validate:"max=255"`
	TTL         string      `json:"ttl" description:"Token TTL in seconds" validate:"required"`
	Permissions Permissions `json:"permissions"`
}

type GetAPITokenInput struct {
	ID string `param:"id" description:"Token ID" validate:"required"`
}

type DeleteAPITokenInput struct {
	ID string `param:"id" description:"Token ID" validate:"required"`
}

type ListAPITokensOutput struct {
	Result []APIToken `json:"result"`
}
