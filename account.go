package figo

// Account is the central domain object of the figo Connect API. The figo Connect API does not
// only consider classical bank accounts as account, but also alternative banking service, e.g.
// credit cards or Paypal.
type Account struct {
	// AccountID is the internal figo Connect account ID
	AccountID string `json:"account_id"`

	// BankID is the internal figo Connect bank ID
	BankID string `json:"bank_id"`

	// Name is the account name
	Name string `json:"name"`

	// Owner is the account owner
	Owner string `json:"owner"`

	// AutoSync indicates whether the account will be automatically synchronized
	AutoSync bool `json:"auto_sync"`

	// AccountNumber national account number
	AccountNumber string `json:"account_number"`

	// BankCode national bank code
	BankCode string `json:"bank_code"`

	// BankName is the bank name
	BankName string `json:"bank_name"`

	// Currency is the three character currency code
	Currency string `json:"currency"`

	// IBAN international bank account number
	IBAN string `json:"iban"`

	// BIC bank identification code
	BIC string `json:"bic"`

	// Type is the account type
	Type string `json:"type"`

	// Icon ist the account icon URL
	Icon string `json:"icon"`

	// SupportedPayments is a map of payment types with it's payment parameters
	SupportedPayments map[string]*SupportedPayment `json:"supported_payments"`

	// SupportedTANSchemes is a list of supported TAN schemes
	SupportedTANSchemes []*SupportedTANSchema `json:"supported_tan_schemes"`

	// PreferredTANScheme is the ID of the TAN scheme preferred by the user
	PreferredTANScheme string `json:"preferred_tan_scheme"`

	// InTotalBalance indicates whether the balance of this account is added to the total balance of accounts
	InTotalBalance bool `json:"in_total_balance"`

	// SavePIN indicates whether the user has chosen to save the PIN on the figo Connect server
	SavePIN bool `json:"save_pin"`

	// Status is the account's synchronization status
	Status *Status `json:"status"`

	// Balance last known account balance
	Balance *Balance `json:"balance"`
}
