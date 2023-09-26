package postgres

import (
	"context"
	"database/sql"
	"errors"

	"github.com/yach36/clean-arch-prac/domain/model"
	"github.com/yach36/clean-arch-prac/domain/repository"
)

type userRepository struct {
	DB *sql.DB
}

var _ repository.IUserRepository = (*userRepository)(nil)

func NewUserRepository(db *sql.DB) *userRepository {
	return &userRepository{
		DB: db,
	}
}

func (r *userRepository) FindAll(ctx context.Context) ([]*model.User, error) {
	cmd := "SELECT * FROM users"
	rows, err := r.DB.Query(cmd)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]*model.User, 0)
	for rows.Next() {
		var u model.User
		err := rows.Scan(&u.ID, &u.Name, &u.Age)
		if err != nil {
			return nil, err
		}
		users = append(users, &u)
	}
	return users, nil
}

func (r *userRepository) Find(ctx context.Context, id int) (*model.User, error) {
	cmd := "SELECT * FROM users where id = $1"
	row := r.DB.QueryRow(cmd, id)
	var user model.User
	err := row.Scan(&user.ID, &user.Name, &user.Age)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("not found")
		}
		return nil, err
	}
	return &user, nil
}
