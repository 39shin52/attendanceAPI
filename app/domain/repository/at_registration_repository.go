package repository

import (
	"context"

	entity "github.com/39shin52/attendanceApp/app/domain/entitiy"
)

type AttendanceRepository interface {
	InsertAttendance(context.Context, *entity.Attendance) error
}
