package patients

// Patient is the public representation of a clinic patient.
type Patient struct {
	ID        string  `json:"id"`
	Code      string  `json:"code"`
	FullName  string  `json:"full_name"`
	Mobile    string  `json:"mobile"`
	Age       int     `json:"age,omitempty"`
	Gender    string  `json:"gender,omitempty"`
	Visits    int     `json:"visits"`
	Invoices  int     `json:"invoices"`
	TotalPaid int64   `json:"total_paid"` // in Rials
	Wallet    int64   `json:"wallet"`     // credit balance in Rials
	Status    string  `json:"status"`
}

// PatientProfile extends Patient with notes for the detail view.
type PatientProfile struct {
	Patient
	Notes string `json:"notes,omitempty"`
}
