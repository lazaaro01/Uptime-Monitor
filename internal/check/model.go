package check

import "time"

type Status string

const (
	StatusUp    Status = "UP"
	StatusDown  Status = "DoWN"
)

type Check struct {
	ID           string
	URL          string
	Status       Status
	LastChecked  time.Time
}