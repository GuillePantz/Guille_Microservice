// appointment_handler.go

package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/GuillePantz/Guille_Microservice/models"
	"github.com/gorilla/mux"
)

// CreateAppointment creates a new appointment
func CreateAppointment(w http.ResponseWriter, r *http.Request, db *database.db) {
	var appointment models.Appointment

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&appointment); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	// Save the appointment to the database (use the database package to interact with the database)
	err := db.SaveAppointment(appointment)
    if err != nil {
        http.Error(w, "Failed to save appointment", http.StatusInternalServerError)
        return
    }

	// Return the created appointment
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(appointment)
}

// GetAppointment gets an appointment by ID
func GetAppointment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	// Retrieve the appointment from the database by ID (use the database package)
	// ...

	// Return the appointment
	// ...
}

// UpdateAppointment updates an appointment by ID
func UpdateAppointment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var updatedAppointment models.Appointment

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&updatedAppointment); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	// Update the appointment in the database by ID (use the database package)
	// ...

	// Return the updated appointment
	// ...
}

// DeleteAppointment deletes an appointment by ID
func DeleteAppointment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	// Delete the appointment from the database by ID (use the database package)
	// ...

	w.WriteHeader(http.StatusNoContent)
}
