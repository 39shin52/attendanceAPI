package usecase

import (
	"context"
	"fmt"
	"time"

	entity "github.com/39shin52/attendanceApp/app/domain/entitiy"
	"github.com/39shin52/attendanceApp/app/domain/repository"
	"github.com/39shin52/attendanceApp/app/domain/repository/transaction"
)

type AttendanceUsecase struct {
	attendanceRepository repository.AttendanceRepository
	userRepository       repository.UserRepository
	tx                   transaction.TxAdmin
}

func NewAttendanceUsecase(attendanceReporitory repository.AttendanceRepository, userReposiory repository.UserRepository, tx transaction.TxAdmin) *AttendanceUsecase {
	return &AttendanceUsecase{attendanceRepository: attendanceReporitory, userRepository: userReposiory, tx: tx}
}

func (au *AttendanceUsecase) RegisterAttendance(ctx context.Context, id string, name string, time time.Time) error {
	_, err := au.userRepository.SelectUserByID(ctx, id)
	if err != nil {
		return fmt.Errorf("user is not found: %v", err)
	}

	newAttendance := entity.ConvertAttendance(id, name, time)

	if err := au.tx.Transaction(ctx, func(ctx context.Context) error {
		if err := au.attendanceRepository.InsertAttendance(ctx, newAttendance); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}
