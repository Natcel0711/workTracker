package api

import (
	"fmt"
	"time"
)

type MySQLDate time.Time

func (d *MySQLDate) Scan(value interface{}) error {
    if value == nil {
        return nil
    }
    t, ok := value.(time.Time)
    if !ok {
        return fmt.Errorf("Failed to scan MySQLDate value")
    }
    *d = MySQLDate(t)
    return nil
}

type Timesheet struct {
    ID           uint64    `json:"id"`
    CompanyName  string    `json:"company_name"`
    HoursWorked  float64   `json:"hours_worked"`
    DateWorked   MySQLDate `json:"date_worked"`
    CreatedAt    time.Time `json:"created_at"`
}