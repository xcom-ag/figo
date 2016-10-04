package figo

type Payment struct {
	PaymentID             string  `json:"payment_id"`
	AccountID             string  `json:"account_id"`
	Type                  string  `json:"type"`
	Name                  string  `json:"name"`
	AccountNumber         string  `json:"account_number"`
	BankCode              string  `json:"bank_code"`
	IBAN                  string  `json:"iban"`
	Amount                float64 `json:"amount"`
	Currency              string  `json:"currency"`
	Purpose               string  `json:"purpose"`
	TextKey               int     `json:"text_key"`
	TextKeyExtension      int     `json:"text_key_extension"`
	NotificationRecipient string  `json:"notification_recipient"`
	Interval              string  `json:"interval"`
	ExecutionDay          int     `json:"execution_day"`
	FirstExecutionDate    string  `json:"first_execution_date"`
	LastExecutionDate     string  `json:"last_execution_date"`
	Cents                 bool    `json:"cents"`
}
