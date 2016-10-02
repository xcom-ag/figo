package figo

// Balance information about account's balance and oconfigured limits
type Balance struct {
	Balance              float64 `json:"balance"`
	BalanceDate          string  `json:"balance_date"`
	CreditLine           float64 `json:"credit_line"`
	MonthlySpendingLimit float64 `json:"monthly_spending_limit"`
	Status               *Status `json:"status"`
}
