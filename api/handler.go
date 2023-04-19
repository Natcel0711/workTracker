package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Natcel0711/workTracker/db"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Healthy")
}

func HandleSomeRequest(w http.ResponseWriter, r *http.Request) {
	// Get the database connection
	dbConn := db.GetDB()

	// Use the database connection to perform database operations
	// For example, you can query the database
	rows, err := dbConn.Query("SELECT * FROM timesheet")
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to query database: %v", err), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Process the query results and write response to the client
	// ...
}

func GetTimesheet(w http.ResponseWriter, r *http.Request) {
	// Get the database connection
	dbConn := db.GetDB()

	// Query timesheet entries from the database
	rows, err := dbConn.Query("SELECT id, company_name, hours_worked, date_worked, created_at FROM timesheet")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Iterate through the rows and populate Timesheet objects
	var timesheets []Timesheet
	for rows.Next() {
		var timesheet Timesheet
		err := rows.Scan(&timesheet.ID, &timesheet.CompanyName, &timesheet.HoursWorked, &timesheet.DateWorked, &timesheet.CreatedAt)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		timesheets = append(timesheets, timesheet)
	}

	// Check for any errors during iteration
	err = rows.Err()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the timesheets as JSON response
	// Note: You can customize the response format according to your needs
	// For example, you can use a JSON encoding library like encoding/json to serialize the timesheets
	jsonBytes, err := json.Marshal(timesheets)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	jsonString := string(jsonBytes)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(jsonString))
}

func CreateTimesheet(w http.ResponseWriter, r *http.Request) {
	var timesheet TimesheetCreate

	err := json.NewDecoder(r.Body).Decode(&timesheet)
	if err != nil {
		http.Error(w, "Failed to parse request body"+"\n"+err.Error(), http.StatusBadRequest)
		return
	}

	dateWorked, err := time.Parse("2006-01-02", timesheet.DateWorked)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Failed to parse date: %v", err)
		return
	}
	timesheet.CreatedAt = time.Now()
	dbConn := db.GetDB()

	result, err := dbConn.Exec("INSERT INTO timesheet (company_name, hours_worked, date_worked, created_at) VALUES (?, ?, ?, ?)",
		timesheet.CompanyName, timesheet.HoursWorked, dateWorked, timesheet.CreatedAt)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Failed to insert timesheet: %v", err)
		return
	}
	// Get inserted row ID
	id, err := result.LastInsertId()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Failed to get last inserted ID: %v", err)
		return
	}
	timesheet.ID = uint64(id)
	response, err := json.Marshal(timesheet)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Failed to marshal response: %v", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func UpdateTimesheet(w http.ResponseWriter, r *http.Request) {
	// Decode JSON request into Timesheet struct
	var timesheet TimesheetCreate
	err := json.NewDecoder(r.Body).Decode(&timesheet)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Failed to decode request body: %v", err)
		return
	}

	// Convert date string to time.Time
	dateWorked, err := time.Parse("2006-01-02", timesheet.DateWorked)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Failed to parse date: %v", err)
		return
	}
	dbConn := db.GetDB()
	// Update timesheet in MySQL table
	result, err := dbConn.Exec("UPDATE timesheet SET company_name=?, hours_worked=?, date_worked=? WHERE id=?",
		timesheet.CompanyName, timesheet.HoursWorked, dateWorked, timesheet.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Failed to update timesheet: %v", err)
		return
	}

	// Get number of rows affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Failed to get rows affected: %v", err)
		return
	}

	if rowsAffected == 0 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Timesheet not found")
		return
	}

	// Convert timesheet struct to JSON
	response, err := json.Marshal(timesheet)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Failed to marshal response: %v", err)
		return
	}

	// Set Content-Type header and write response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func DeleteTimesheet(w http.ResponseWriter, r *http.Request) {
	// Parse ID from URL parameters
	id := r.FormValue("id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "ID is required")
		return
	}

	dbConn := db.GetDB()

	// Delete timesheet from MySQL table
	result, err := dbConn.Exec("DELETE FROM timesheet WHERE id=?", id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Failed to delete timesheet: %v", err)
		return
	}

	// Get number of rows affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Failed to get rows affected: %v", err)
		return
	}

	if rowsAffected == 0 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Timesheet not found")
		return
	}

	// Set Content-Type header and write response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Timesheet with ID %s has been deleted", id)
}
