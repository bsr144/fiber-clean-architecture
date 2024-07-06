package requests

type PatientFhir struct {
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
