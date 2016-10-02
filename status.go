package figo

// Status holds information about the account's synchronization status
type Status struct {

	// Code is the internal figo connect status code
	Code int `json:"code"`

	// Message is the human readable error message
	Message string `json:"message"`

	// SyncTimestamp of last synchronization
	SyncTimestamp string `json:"sync_timestamp"`

	// SuccessTimestamp of last successful synchronization
	SuccessTimestamp string `json:"success_timestamp"`
}
