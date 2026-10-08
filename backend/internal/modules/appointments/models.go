package appointments

// Request represents an inbound appointment or deposit request.
type Request struct {
	ID          string `json:"id"`
	Channel     string `json:"channel"` // telegram | phone | instagram | whatsapp
	Type        string `json:"type"`    // reserve | deposit
	Status      string `json:"status"`  // pending | accepted | rejected
	PatientName string `json:"patient_name"`
	Mobile      string `json:"mobile,omitempty"`
	Service     string `json:"service,omitempty"`
	PreferredAt string `json:"preferred_at,omitempty"` // ISO or human text
	Amount      int64  `json:"amount,omitempty"`      // for deposit, in Rials
	Notes       string `json:"notes,omitempty"`
	CreatedAt   string `json:"created_at"`
	DecidedAt   string `json:"decided_at,omitempty"`
	DecidedBy   string `json:"decided_by,omitempty"`
}

// DecideRequest is the body for accept/reject.
type DecideRequest struct {
	Action string `json:"action"` // accept | reject
	Notes  string `json:"notes,omitempty"`
}
