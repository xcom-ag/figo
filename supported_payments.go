package figo

type SupportedPayment struct {
	AllowedRecipients []string `json:"allowed_recipients"`
	MaxPurposeLength  int      `json:"max_purpose_length"`
	SupportedTextKeys []int    `json:"supported_text_keys"`
	MinScheduledDate  string   `json:"min_scheduled_date"`
	MaxScheduledDate  string   `json:"max_scheduled_date"`
}
