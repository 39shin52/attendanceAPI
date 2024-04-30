package usecase

import (
	"context"
	"fmt"

	entity "github.com/39shin52/attendanceApp/app/domain/entitiy"
	"github.com/39shin52/attendanceApp/app/domain/repository"
	"github.com/39shin52/attendanceApp/app/domain/repository/transaction"
	"github.com/google/uuid"
)

type UserUsecase struct {
	userRepository repository.UserRepository
	tx             transaction.TxAdmin
}

func NewUserUsecase(userRepository repository.UserRepository, tx transaction.TxAdmin) *UserUsecase {
	return &UserUsecase{
		userRepository: userRepository,
		tx:             tx,
	}
}

func (uu *UserUsecase) CreateUser(ctx context.Context, name string, affiliation string) (*entity.User, error) {
	id := uuid.NewString()

	newUser := &entity.User{
		ID:          id,
		Name:        name,
		Affiliation: affiliation,
	}

	if err := uu.tx.Transaction(ctx, func(ctx context.Context) error {
		user, err := uu.userRepository.InsertUser(ctx, newUser)
		if err != nil {
			return err
		}

		if user != newUser {
			return fmt.Errorf("some error occured")
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return newUser, nil
}
