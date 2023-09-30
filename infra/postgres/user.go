package postgres

import (
	"context"
	"database/sql"

	"github.com/yach36/clean-arch-prac/domain/model"
	"github.com/yach36/clean-arch-prac/domain/repository"
	"github.com/yach36/clean-arch-prac/utils/cerrors"
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
			return nil, cerrors.NotFound.New("record not found")
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Add(ctx context.Context, user *model.User) error {
	cmd := "INSERT INTO users (name, age) VALUES($1, $2)"
	_, err := r.DB.Exec(cmd, user.Name, user.Age)
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepository) Delete(ctx context.Context, id int) error {
	cmd := "DELETE FROM users WHERE id = $1"
	_, err := r.DB.Exec(cmd, id)
	if err != nil {
		return err
	}
	return nil
}
