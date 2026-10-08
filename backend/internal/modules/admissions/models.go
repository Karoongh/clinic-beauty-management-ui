package admissions

// Admission represents a patient acceptance record.
type Admission struct {
	ID          string `json:"id"`
	Number      int    `json:"number"` // starts from 1 per clinic policy
	Date        string `json:"date"`   // YYYY-MM-DD
	PatientID   string `json:"patient_id"`
	PatientName string `json:"patient_name"`
	Channel     string `json:"channel"` // walk-in | phone | telegram | instagram | whatsapp
	Service     string `json:"service,omitempty"`
	Doctor      string `json:"doctor,omitempty"`
	Notes       string `json:"notes,omitempty"`
	CreatedAt   string `json:"created_at"`
}

// CreateAdmissionRequest is the body for creating a new admission.
type CreateAdmissionRequest struct {
	PatientID string `json:"patient_id"`
	Channel   string `json:"channel"`
	Service   string `json:"service,omitempty"`
	Doctor    string `json:"doctor,omitempty"`
	Notes     string `json:"notes,omitempty"`
	Date      string `json:"date,omitempty"` // optional, defaults to today
}
