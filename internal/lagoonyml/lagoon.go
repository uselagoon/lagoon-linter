package lagoonyml

import "encoding/json"

// Ingress represents a Lagoon route.
type Ingress struct {
	TLSACME     string            `json:"tls-acme"`
	Insecure    string            `json:"insecure"`
	HSTS        string            `json:"hsts"`
	Annotations map[string]string `json:"annotations"`
}

// LagoonRoute can be either a string or a map[string]Ingress, so we must
// implement a custom unmarshaller.
type LagoonRoute struct {
	Name      string
	Ingresses map[string]Ingress
}

// UnmarshalJSON implements json.Unmarshaler.
func (lr *LagoonRoute) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &lr.Name); err == nil {
		return nil
	}
	return json.Unmarshal(data, &lr.Ingresses)
}

// Environment represents a Lagoon environment.
type Environment struct {
	Routes []map[string][]LagoonRoute `json:"routes"`
}

// Lagoon represents the .lagoon.yml file.
type Lagoon struct {
	Environments map[string]Environment `json:"environments"`
}
