package types

type HealthcheckTLSDefinition struct {
	Target          string `json:"target" validate:"required"`
	Port            uint   `json:"port" validate:"required,max=65535,min=1"`
	Key             string `json:"key,omitempty"`
	Cert            string `json:"cert,omitempty"`
	Cacert          string `json:"cacert,omitempty"`
	ServerName      string `json:"server-name,omitempty"`
	Insecure        bool   `json:"insecure"`
	ExpirationDelay string `json:"expiration-delay"`
}

type CreateTLSHealthcheckInput struct {
	Name        string            `json:"name" description:"Healthcheck name" validate:"required,max=255,min=1"`
	Description string            `json:"description" description:"Healthcheck description" validate:"max=255"`
	Labels      map[string]string `json:"labels" description:"Healthcheck labels" validate:"dive,keys,max=255,min=1,endkeys,max=255,min=1"`
	Interval    string            `json:"interval" description:"Healthcheck interval" validate:"required"`
	Enabled     bool              `json:"bool" description:"Enable the healthcheck on the appclacks platform"`
	Timeout     string            `json:"timeout" validate:"required"`
	HealthcheckTLSDefinition
}

type UpdateTLSHealthcheckInput struct {
	ID          string            `param:"id" description:"Healthcheck ID" validate:"required,uuid"`
	Name        string            `json:"name" description:"Healthcheck name" validate:"required,max=255,min=1"`
	Description string            `json:"description" description:"Healthcheck description" validate:"max=255"`
	Labels      map[string]string `json:"labels" description:"Healthcheck labels" validate:"dive,keys,max=255,min=1,endkeys,max=255,min=1"`
	Interval    string            `json:"interval" description:"Healthcheck interval" validate:"required"`
	Timeout     string            `json:"timeout" validate:"required"`
	Enabled     bool              `json:"enabled" description:"Enable the healthcheck on the appclacks platform"`
	HealthcheckTLSDefinition
}
