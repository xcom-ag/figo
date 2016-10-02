package figo

// User informations about the figo Connect user
type User struct {
	// UserID is the internal figo Connect user ID
	UserID string `json:"user_id"`

	// Name holds the user's first and last name
	Name string `json:"name"`

	// Email address
	Email string `json:"email"`

	// Address is the postal address for bills, etc.
	Address *Address `json:"address"`

	// VerifiedEmail indicates whether the email address has been verified
	VerifiedEmail bool `json:"verified_email"`

	// SendNewsLetter indicates whether the user has agreed to be contacted by email.
	SendNewsletter bool `json:"send_newsletter"`

	// Language is the two letter code of the user's preferred language
	Language string `json:"language"`

	// JoinDate is the timestamp of the figo account registration
	JoinDate string `json:"join_date"`

	// ForceReset indicates whether all local data must be cleared from the device and refetched from the figo Connect server
	ForceReset bool `json:"force_reset"`

	// Filters is an array of filters defined by the user
	Filters []string `json:"filters"`
}
