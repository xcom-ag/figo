package figo

// Address holds information about a postal address
type Address struct {
	City       string `json:"city"`
	Country    string `json:"country"`
	Company    string `json:"company"`
	Street     string `json:"street"`
	PostalCode string `json:"postal_code"`
	Street2    string `json:"street2"`
}
