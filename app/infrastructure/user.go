package infrastructure

import (
	"context"
	"database/sql"

	entity "github.com/39shin52/attendanceApp/app/domain/entitiy"
	"github.com/39shin52/attendanceApp/app/domain/repository"
)

type userRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) repository.UserRepository {
	return &userRepositoryImpl{db: db}
}

func (ur *userRepositoryImpl) SelectUserByID(ctx context.Context, id string) (*entity.User, error) {
	query := `select id, name, affiliation from user where id = ?`

	row := ur.db.QueryRow(query, id)

	return convertUser(row)
}

func convertUser(row *sql.Row) (*entity.User, error) {
	user := new(entity.User)
	if err := row.Scan(&user.ID, &user.Name, &user.Affiliation); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return user, nil
}
