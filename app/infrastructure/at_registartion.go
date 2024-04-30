package infrastructure

import (
	"context"
	"database/sql"

	entity "github.com/39shin52/attendanceApp/app/domain/entitiy"
	"github.com/39shin52/attendanceApp/app/domain/repository"
)

type attendanceRepositoryImpl struct {
	db *sql.DB
}

func NewAttendanceRepository(db *sql.DB) repository.AttendanceRepository {
	return &attendanceRepositoryImpl{db: db}
}

func (ar *attendanceRepositoryImpl) InsertAttendance(ctx context.Context, attendance *entity.Attendance) error {
	query := `insert into attendance(id, time) values (?, ?)`

	_, err := ar.db.ExecContext(ctx, query, attendance.ID, attendance.Time)
	if err != nil {
		return err
	}

	return nil
}
