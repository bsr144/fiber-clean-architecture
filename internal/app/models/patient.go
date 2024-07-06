package models

type Patient struct {
	ID           string         `json:"id"`
	IDd          string         `json:"_id"`
	ResourceType string         `json:"resourceType"`
	Active       bool           `json:"active"`
	Name         []HumanName    `json:"name"`
	Telecom      []ContactPoint `json:"telecom"`
}

type HumanName struct {
	Use    string   `json:"use"`
	Family string   `json:"family"`
	Given  []string `json:"given"`
}

type ContactPoint struct {
	System string `json:"system"`
	Value  string `json:"value"`
	Use    string `json:"use"`
}
