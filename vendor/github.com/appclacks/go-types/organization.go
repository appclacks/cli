package types

import (
	"time"
)

type CreateOrganizationOrg struct {
	Name        string `json:"name" description:"Organization name" validate:"required,max=255"`
	Description string `json:"description" description:"Organization description" validate:"max=255"`
}

type CreateOrganizationAccount struct {
	FirstName string `json:"first-name" description:"User first name" validate:"required,max=255"`
	LastName  string `json:"last-name" description:"User last name" validate:"required,max=255"`
	Password  string `json:"password" description:"User password" validate:"required,min=10,max=255"`
	Email     string `json:"email" description:"User email" validate:"required,max=255"`
}

type CreateOrganizationInput struct {
	Organization CreateOrganizationOrg     `json:"organization" validate:"required"`
	Account      CreateOrganizationAccount `json:"account" validate:"required"`
}

type CreateOrganizationOutput struct {
	Organization Organization `json:"organization"`
	Account      Account      `json:"account"`
}

type Organization struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description"`
	CreatedAt   time.Time `json:"created-at"`
}

type Account struct {
	ID        string    `json:"id"`
	FirstName string    `json:"first-name"`
	LastName  string    `json:"last-name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created-at"`
}

type GetOrganizationInput struct {
	ID string `param:"id" description:"Organization ID" validate:"required"`
}

type GetOrganizationOutput struct {
	Organization Organization `json:"organization"`
}
