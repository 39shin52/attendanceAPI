package repository

import (
	"context"

	entity "github.com/39shin52/attendanceApp/app/domain/entitiy"
)

type UserRepository interface {
	SelectUserByID(context.Context, string) (*entity.User, error)
	InsertUser(context.Context, *entity.User) (*entity.User, error)
}
