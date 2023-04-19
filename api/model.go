package api

import (
	"fmt"
	"time"
)

type Timesheet struct {
	ID          uint64         `db:"id" json:"id"`
	CompanyName string         `db:"company_name" json:"company_name"`
	HoursWorked float64        `db:"hours_worked" json:"hours_worked"`
	Incident    string         `db:"incident" json:"incident"`
	Resolved    bool           `db:"resolved" json:"resolved"`
	DateWorked  MySQLDate      `db:"date_worked" json:"date_worked"`
	CreatedAt   MySQLTimestamp `db:"created_at" json:"created_at"`
}

type TimesheetCreate struct {
	ID          uint64    `json:"id,omitempty"`
	CompanyName string    `json:"company_name"`
	HoursWorked float64   `json:"hours_worked"`
	DateWorked  string    `json:"date_worked"`
	Incident    string    `db:"incident"`
	Resolved    bool      `db:"resolved"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
}

type MySQLDate struct {
	time.Time
}

type MySQLTimestamp struct {
	time.Time
}

func (t *MySQLTimestamp) Scan(value interface{}) error {
	switch v := value.(type) {
	case []byte:
		str := string(v)
		tm, err := time.Parse("2006-01-02 15:04:05", str)
		if err != nil {
			return err
		}
		t.Time = tm
	case string:
		tm, err := time.Parse("2006-01-02 15:04:05", v)
		if err != nil {
			return err
		}
		t.Time = tm
	default:
		return fmt.Errorf("unsupported Scan, storing driver.Value type %T into type *MySQLTime", value)
	}
	return nil
}

func (t *MySQLDate) Scan(value interface{}) error {
	switch v := value.(type) {
	case []byte:
		str := string(v)
		tm, err := time.Parse("2006-01-02", str)
		if err != nil {
			return err
		}
		t.Time = tm
	case string:
		tm, err := time.Parse("2006-01-02", v)
		if err != nil {
			return err
		}
		t.Time = tm
	default:
		return fmt.Errorf("unsupported Scan, storing driver.Value type %T into type *MySQLTime", value)
	}
	return nil
}
