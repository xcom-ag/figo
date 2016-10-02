package figo

type SupportedTANSchema struct {
	TANSchemeID string `json:"tan_scheme_id"`
	Name        string `json:"name"`
	MediumName  string `json:"medium_name"`
}
