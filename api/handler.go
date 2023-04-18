package api

import (
	"encoding/json"
	"fmt"
	"net/http"

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

func HandleTimesheetGet(w http.ResponseWriter, r *http.Request) {
	// Get the database connection
	dbConn := db.GetDB()

	// Query timesheet entries from the database
	rows, err := dbConn.Query("SELECT id, company_name, hours_worked FROM timesheet")
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to query timesheet entries from database: %v", err), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Iterate through the query results and build timesheet entries
	var timesheets []Timesheet
	for rows.Next() {
		var timesheet Timesheet
		err := rows.Scan(&timesheet.ID, &timesheet.CompanyName, &timesheet.HoursWorked)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to scan timesheet entry from query result: %v", err), http.StatusInternalServerError)
			return
		}
		timesheets = append(timesheets, timesheet)
	}

	// Write the response back to the client
	json.NewEncoder(w).Encode(timesheets)
}
