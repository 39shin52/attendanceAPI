package request

import "time"

type AttendanceRegistrationRequest struct {
	ID   string
	Name string
	Time time.Time
}
