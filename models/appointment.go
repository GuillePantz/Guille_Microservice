// appointment.go

package models

import "github.com/google/uuid"

// Appointment represents an appointment entity
type Appointment struct {
	ID        uuid.UUID `json:"id"`
	PatientID uuid.UUID `json:"patient_id"`
	// Add other fields as needed
}